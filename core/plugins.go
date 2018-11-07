package core

import (
	"fmt"
	"plugin"

	"github.com/pkg/errors"
)

type Plugins struct {
	ed    *Editor
	plugs []*Plug
	added map[string]bool
}

func NewPlugins(ed *Editor) *Plugins {
	return &Plugins{ed: ed, added: map[string]bool{}}
}

func (p *Plugins) AddPath(path string) {
	if p.added[path] {
		return
	}
	p.added[path] = true

	oplugin, err := plugin.Open(path)
	if err != nil {
		err2 := errors.Wrap(err, fmt.Sprintf("plugin: %v", path))
		p.ed.Error(err2)
		return
	}

	plug := &Plug{Plugin: oplugin, Path: path}
	p.plugs = append(p.plugs, plug)

	p.runOnLoad(plug)
}

//----------

func (p *Plugins) runOnLoad(plug *Plug) {
	// plugin should have this symbol
	f, err := plug.Plugin.Lookup("OnLoad")
	if err != nil {
		// silent error
		return
	}
	// the symbol must implement this signature
	f2, ok := f.(func(*Editor))
	if !ok {
		err := fmt.Errorf("plugin: %v: bad func signature", plug.Path)
		p.ed.Error(err)
		return
	}
	// run onload
	f2(p.ed)
}

//----------

type Plug struct {
	Path   string
	Plugin *plugin.Plugin
}