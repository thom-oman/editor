package xdriver

import (
	"fmt"
	"image"
	"image/draw"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/shm"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xcursor"
	"github.com/jmigpin/editor/driver/xdriver/copypaste"
	"github.com/jmigpin/editor/driver/xdriver/dragndrop"
	"github.com/jmigpin/editor/driver/xdriver/wimage"
	"github.com/jmigpin/editor/driver/xdriver/wmprotocols"
	"github.com/jmigpin/editor/driver/xdriver/xcursors"
	"github.com/jmigpin/editor/driver/xdriver/xinput"
	"github.com/jmigpin/editor/driver/xdriver/xutil"
	"github.com/jmigpin/editor/util/uiutil/event"
)

type Window struct {
	Conn   *xgb.Conn
	Window xproto.Window
	Screen *xproto.ScreenInfo
	GCtx   xproto.Gcontext

	Paste   *copypaste.Paste
	Copy    *copypaste.Copy
	Cursors *xcursors.Cursors
	XInput  *xinput.XInput
	Wmp     *wmprotocols.WMP
	Dnd     *dragndrop.Dnd

	WImg wimage.WImage

	events    chan interface{}
	closeOnce sync.Once
}

func NewWindow() (*Window, error) {
	display := os.Getenv("DISPLAY")

	// help get a display target
	origDisplay := display
	if display == "" {
		switch runtime.GOOS {
		case "windows":
			display = "127.0.0.1:0.0"
		}
	}

	conn, err := xgb.NewConnDisplay(display)
	if err != nil {
		// improve error with hint
		if origDisplay == "" {
			err = fmt.Errorf("%w (Hint: is x11 running?)", err)
		}
		return nil, fmt.Errorf("x11 conn: %w", err)
	}

	// initialize extensions early to avoid concurrent map read/write (XGB issue)
	wimage.Init(conn)

	win := &Window{
		Conn:   conn,
		events: make(chan interface{}, 8),
	}

	if err := win.initialize(); err != nil {
		_ = win.Close() // best effort to close since it was opened
		return nil, fmt.Errorf("win init: %w", err)
	}

	go win.eventLoop()

	return win, nil
}
func (win *Window) initialize() error {
	// Disable xgb logger that prints to stderr
	//xgb.Logger = log.New(ioutil.Discard, "", 0)

	si := xproto.Setup(win.Conn)
	win.Screen = si.DefaultScreen(win.Conn)

	window, err := xproto.NewWindowId(win.Conn)
	if err != nil {
		return err
	}
	win.Window = window

	// event mask
	var evMask uint32 = 0 |
		xproto.EventMaskStructureNotify |
		xproto.EventMaskExposure |
		xproto.EventMaskPropertyChange |
		//xproto.EventMaskPointerMotionHint |
		//xproto.EventMaskButtonMotion |
		xproto.EventMaskPointerMotion |
		xproto.EventMaskButtonPress |
		xproto.EventMaskButtonRelease |
		xproto.EventMaskKeyPress |
		xproto.EventMaskKeyRelease |
		0
	// mask/values order is defined by the protocol
	mask := uint32(xproto.CwEventMask)
	values := []uint32{evMask}

	_ = xproto.CreateWindow(
		win.Conn,
		win.Screen.RootDepth,
		win.Window,
		win.Screen.Root,
		0, 0, 500, 500,
		0, // border width
		xproto.WindowClassInputOutput,
		win.Screen.RootVisual,
		mask, values)

	_ = xproto.MapWindow(win.Conn, window)

	if err := xutil.LoadAtoms(win.Conn, &Atoms, false); err != nil {
		return err
	}

	// graphical context
	gCtx, err := xproto.NewGcontextId(win.Conn)
	if err != nil {
		return err
	}
	win.GCtx = gCtx

	gmask := uint32(0)
	gvalues := []uint32{}
	c2 := xproto.CreateGCChecked(win.Conn, win.GCtx, xproto.Drawable(win.Window), gmask, gvalues)
	if err := c2.Check(); err != nil {
		return err
	}

	xi, err := xinput.NewXInput(win.Conn)
	if err != nil {
		return err
	}
	win.XInput = xi

	dnd, err := dragndrop.NewDnd(win.Conn, win.Window)
	if err != nil {
		return err
	}
	win.Dnd = dnd

	paste, err := copypaste.NewPaste(win.Conn, win.Window)
	if err != nil {
		return err
	}
	win.Paste = paste

	copy, err := copypaste.NewCopy(win.Conn, win.Window)
	if err != nil {
		return err
	}
	win.Copy = copy

	c, err := xcursors.NewCursors(win.Conn, win.Window)
	if err != nil {
		return err
	}
	win.Cursors = c

	opt := &wimage.Options{win.Conn, win.Window, win.Screen, win.GCtx}
	img, err := wimage.NewWImage(opt)
	if err != nil {
		return err
	}
	win.WImg = img

	wmp, err := wmprotocols.NewWMP(win.Conn, win.Window)
	if err != nil {
		return err
	}
	win.Wmp = wmp

	return nil
}

func (win *Window) Close() error {
	win.closeOnce.Do(func() {
		if win.WImg != nil { // might not be instantiated if failing at init
			// closing before the connection to prevent errors like "sending on a closed channel" from xgb
			_ = win.WImg.Close()
		}
		win.Conn.Close() // conn.WaitForEvent() will return with (nil,nil)
	})
	return nil
}

func (win *Window) NextEvent() interface{} {
	return <-win.events
}

func (win *Window) eventLoop() {
	// TODO: improve (maybe only option is window2 interface)
	// unblocks waiting for wimg.putimagecompleted
	defer func() {
		if u, ok := win.WImg.(*wimage.ShmWImage); ok {
			u.PutImageCompleted()
		}
	}()

	for {
		ok := win.handleEvent(win.events)
		if !ok {
			break
		}
	}
}

func (win *Window) handleEvent(events chan interface{}) bool {
	ev, xerr := win.Conn.WaitForEvent()
	if ev == nil && xerr == nil {
		events <- &event.WindowClose{}
		return false
	}
	if xerr != nil {
		events <- error(xerr)
	}
	if ev != nil {
		switch t := ev.(type) {
		case xproto.ConfigureNotifyEvent: // structure (position,size,...)
			//x, y := int(t.X), int(t.Y) // commented: must use (0,0)
			w, h := int(t.Width), int(t.Height)
			r := image.Rect(0, 0, w, h)
			events <- &event.WindowResize{Rect: r}
		case xproto.ExposeEvent: // region needs paint
			//x, y := int(t.X), int(t.Y) // commented: must use (0,0)
			w, h := int(t.Width), int(t.Height)
			r := image.Rect(0, 0, w, h)
			events <- &event.WindowExpose{Rect: r}
		case xproto.MapNotifyEvent: // window mapped (created)
		case xproto.ReparentNotifyEvent: // window rerooted

		case shm.CompletionEvent:
			win.WImg.PutImageCompleted()
		case xproto.MappingNotifyEvent: // keyboard mapping
			if err := win.XInput.ReadMapTable(); err != nil {
				events <- err
			}

		case xproto.KeyPressEvent:
			events <- win.XInput.KeyPress(&t)
		case xproto.KeyReleaseEvent:
			events <- win.XInput.KeyRelease(&t)
		case xproto.ButtonPressEvent:
			events <- win.XInput.ButtonPress(&t)
		case xproto.ButtonReleaseEvent:
			events <- win.XInput.ButtonRelease(&t)
		case xproto.MotionNotifyEvent:
			events <- win.XInput.MotionNotify(&t)

		case xproto.SelectionNotifyEvent:
			win.Paste.OnSelectionNotify(&t)
			win.Dnd.OnSelectionNotify(&t)
		case xproto.SelectionRequestEvent:
			if err := win.Copy.OnSelectionRequest(&t); err != nil {
				events <- err
			}
		case xproto.SelectionClearEvent:
			win.Copy.OnSelectionClear(&t)

		case xproto.ClientMessageEvent:
			deleteWindow := win.Wmp.OnClientMessage(&t)
			if deleteWindow {
				events <- &event.WindowClose{}
				return false
			}
			if ev2, err, ok := win.Dnd.OnClientMessage(&t); ok {
				if err != nil {
					events <- err
				} else {
					events <- ev2
				}
			}

		case xproto.PropertyNotifyEvent:
			win.Paste.OnPropertyNotify(&t)

		default:
			log.Printf("unhandled event: %#v", ev)
		}
	}
	return true
}

func (win *Window) SetWindowName(str string) {
	b := []byte(str)
	_ = xproto.ChangeProperty(
		win.Conn,
		xproto.PropModeReplace,
		win.Window,       // requestor window
		Atoms.NetWMName,  // property
		Atoms.Utf8String, // target
		8,                // format
		uint32(len(b)),
		b)
}

//func (win *Window) GetGeometry() (*xproto.GetGeometryReply, error) {
//	drawable := xproto.Drawable(win.Window)
//	cookie := xproto.GetGeometry(win.Conn, drawable)
//	return cookie.Reply()
//}

func (win *Window) Image() draw.Image {
	return win.WImg.Image()
}
func (win *Window) PutImage(rect image.Rectangle) error {
	return win.WImg.PutImage(rect)
}
func (win *Window) ResizeImage(r image.Rectangle) error {
	ib := win.Image().Bounds()
	if !r.Eq(ib) {
		err := win.WImg.Resize(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (win *Window) WarpPointer(p image.Point) {
	// warp pointer only if the window has input focus
	cookie := xproto.GetInputFocus(win.Conn)
	reply, err := cookie.Reply()
	if err != nil {
		log.Print(err)
		return
	}
	if reply.Focus != win.Window {
		return
	}
	_ = xproto.WarpPointer(
		win.Conn,
		xproto.WindowNone,
		win.Window,
		0, 0, 0, 0,
		int16(p.X), int16(p.Y))
}

func (win *Window) QueryPointer() (image.Point, error) {
	cookie := xproto.QueryPointer(win.Conn, win.Window)
	r, err := cookie.Reply()
	if err != nil {
		return image.ZP, err
	}
	p := image.Point{int(r.WinX), int(r.WinY)}
	return p, nil
}

func (win *Window) GetCPPaste(i event.CopyPasteIndex, fn func(string, error)) {
	win.Paste.Get(i, fn)
}
func (win *Window) SetCPCopy(i event.CopyPasteIndex, v string) error {
	return win.Copy.Set(i, v)
}

func (win *Window) SetCursor(c event.Cursor) {
	sc := func(c2 xcursors.Cursor) {
		err := win.Cursors.SetCursor(c2)
		if err != nil {
			log.Print(err)
		}
	}
	switch c {
	case event.NoneCursor:
		sc(xcursors.XCNone)
	case event.DefaultCursor:
		sc(xcursors.XCNone)
	case event.NSResizeCursor:
		sc(xcursor.SBVDoubleArrow)
	case event.WEResizeCursor:
		sc(xcursor.SBHDoubleArrow)
	case event.CloseCursor:
		sc(xcursor.XCursor)
	case event.MoveCursor:
		sc(xcursor.Fleur)
	case event.PointerCursor:
		sc(xcursor.Hand2)
	case event.BeamCursor:
		sc(xcursor.XTerm)
	case event.WaitCursor:
		sc(xcursor.Watch)
	}
}

//----------

var Atoms struct {
	NetWMName  xproto.Atom `loadAtoms:"_NET_WM_NAME"`
	Utf8String xproto.Atom `loadAtoms:"UTF8_STRING"`
}
