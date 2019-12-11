// Code generated by "stringer -type=_wm -output zwm.go"; DO NOT EDIT.

package windriver

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[_WM_NULL-0]
	_ = x[_WM_CREATE-1]
	_ = x[_WM_DESTROY-2]
	_ = x[_WM_MOVE-3]
	_ = x[_WM_SIZE-5]
	_ = x[_WM_ACTIVATE-6]
	_ = x[_WM_SETFOCUS-7]
	_ = x[_WM_KILLFOCUS-8]
	_ = x[_WM_ENABLE-10]
	_ = x[_WM_SETREDRAW-11]
	_ = x[_WM_SETTEXT-12]
	_ = x[_WM_GETTEXT-13]
	_ = x[_WM_GETTEXTLENGTH-14]
	_ = x[_WM_PAINT-15]
	_ = x[_WM_CLOSE-16]
	_ = x[_WM_QUERYENDSESSION-17]
	_ = x[_WM_QUIT-18]
	_ = x[_WM_QUERYOPEN-19]
	_ = x[_WM_ERASEBKGND-20]
	_ = x[_WM_SYSCOLORCHANGE-21]
	_ = x[_WM_ENDSESSION-22]
	_ = x[_WM_SYSTEMERROR-23]
	_ = x[_WM_SHOWWINDOW-24]
	_ = x[_WM_CTLCOLOR-25]
	_ = x[_WM_WININICHANGE-26]
	_ = x[_WM_SETTINGCHANGE-26]
	_ = x[_WM_DEVMODECHANGE-27]
	_ = x[_WM_ACTIVATEAPP-28]
	_ = x[_WM_FONTCHANGE-29]
	_ = x[_WM_TIMECHANGE-30]
	_ = x[_WM_CANCELMODE-31]
	_ = x[_WM_SETCURSOR-32]
	_ = x[_WM_MOUSEACTIVATE-33]
	_ = x[_WM_CHILDACTIVATE-34]
	_ = x[_WM_QUEUESYNC-35]
	_ = x[_WM_GETMINMAXINFO-36]
	_ = x[_WM_PAINTICON-38]
	_ = x[_WM_ICONERASEBKGND-39]
	_ = x[_WM_NEXTDLGCTL-40]
	_ = x[_WM_SPOOLERSTATUS-42]
	_ = x[_WM_DRAWITEM-43]
	_ = x[_WM_MEASUREITEM-44]
	_ = x[_WM_DELETEITEM-45]
	_ = x[_WM_VKEYTOITEM-46]
	_ = x[_WM_CHARTOITEM-47]
	_ = x[_WM_SETFONT-48]
	_ = x[_WM_GETFONT-49]
	_ = x[_WM_SETHOTKEY-50]
	_ = x[_WM_GETHOTKEY-51]
	_ = x[_WM_QUERYDRAGICON-55]
	_ = x[_WM_COMPAREITEM-57]
	_ = x[_WM_COMPACTING-65]
	_ = x[_WM_WINDOWPOSCHANGING-70]
	_ = x[_WM_WINDOWPOSCHANGED-71]
	_ = x[_WM_POWER-72]
	_ = x[_WM_COPYDATA-74]
	_ = x[_WM_CANCELJOURNAL-75]
	_ = x[_WM_NOTIFY-78]
	_ = x[_WM_INPUTLANGCHANGEREQUEST-80]
	_ = x[_WM_INPUTLANGCHANGE-81]
	_ = x[_WM_TCARD-82]
	_ = x[_WM_HELP-83]
	_ = x[_WM_USERCHANGED-84]
	_ = x[_WM_NOTIFYFORMAT-85]
	_ = x[_WM_CONTEXTMENU-123]
	_ = x[_WM_STYLECHANGING-124]
	_ = x[_WM_STYLECHANGED-125]
	_ = x[_WM_DISPLAYCHANGE-126]
	_ = x[_WM_GETICON-127]
	_ = x[_WM_SETICON-128]
	_ = x[_WM_NCCREATE-129]
	_ = x[_WM_NCDESTROY-130]
	_ = x[_WM_NCCALCSIZE-131]
	_ = x[_WM_NCHITTEST-132]
	_ = x[_WM_NCPAINT-133]
	_ = x[_WM_NCACTIVATE-134]
	_ = x[_WM_GETDLGCODE-135]
	_ = x[_WM_NCMOUSEMOVE-160]
	_ = x[_WM_NCLBUTTONDOWN-161]
	_ = x[_WM_NCLBUTTONUP-162]
	_ = x[_WM_NCLBUTTONDBLCLK-163]
	_ = x[_WM_NCRBUTTONDOWN-164]
	_ = x[_WM_NCRBUTTONUP-165]
	_ = x[_WM_NCRBUTTONDBLCLK-166]
	_ = x[_WM_NCMBUTTONDOWN-167]
	_ = x[_WM_NCMBUTTONUP-168]
	_ = x[_WM_NCMBUTTONDBLCLK-169]
	_ = x[_WM_KEYDOWN-256]
	_ = x[_WM_KEYUP-257]
	_ = x[_WM_CHAR-258]
	_ = x[_WM_DEADCHAR-259]
	_ = x[_WM_SYSKEYDOWN-260]
	_ = x[_WM_SYSKEYUP-261]
	_ = x[_WM_SYSCHAR-262]
	_ = x[_WM_SYSDEADCHAR-263]
	_ = x[_WM_KEYLAST-264]
	_ = x[_WM_IME_STARTCOMPOSITION-269]
	_ = x[_WM_IME_ENDCOMPOSITION-270]
	_ = x[_WM_IME_COMPOSITION-271]
	_ = x[_WM_IME_KEYLAST-271]
	_ = x[_WM_INITDIALOG-272]
	_ = x[_WM_COMMAND-273]
	_ = x[_WM_SYSCOMMAND-274]
	_ = x[_WM_TIMER-275]
	_ = x[_WM_HSCROLL-276]
	_ = x[_WM_VSCROLL-277]
	_ = x[_WM_INITMENU-278]
	_ = x[_WM_INITMENUPOPUP-279]
	_ = x[_WM_MENUSELECT-287]
	_ = x[_WM_MENUCHAR-288]
	_ = x[_WM_ENTERIDLE-289]
	_ = x[_WM_CTLCOLORMSGBOX-306]
	_ = x[_WM_CTLCOLOREDIT-307]
	_ = x[_WM_CTLCOLORLISTBOX-308]
	_ = x[_WM_CTLCOLORBTN-309]
	_ = x[_WM_CTLCOLORDLG-310]
	_ = x[_WM_CTLCOLORSCROLLBAR-311]
	_ = x[_WM_CTLCOLORSTATIC-312]
	_ = x[_WM_MOUSEMOVE-512]
	_ = x[_WM_LBUTTONDOWN-513]
	_ = x[_WM_LBUTTONUP-514]
	_ = x[_WM_LBUTTONDBLCLK-515]
	_ = x[_WM_RBUTTONDOWN-516]
	_ = x[_WM_RBUTTONUP-517]
	_ = x[_WM_RBUTTONDBLCLK-518]
	_ = x[_WM_MBUTTONDOWN-519]
	_ = x[_WM_MBUTTONUP-520]
	_ = x[_WM_MBUTTONDBLCLK-521]
	_ = x[_WM_MOUSEWHEEL-522]
	_ = x[_WM_MOUSEHWHEEL-526]
	_ = x[_WM_PARENTNOTIFY-528]
	_ = x[_WM_ENTERMENULOOP-529]
	_ = x[_WM_EXITMENULOOP-530]
	_ = x[_WM_NEXTMENU-531]
	_ = x[_WM_SIZING-532]
	_ = x[_WM_CAPTURECHANGED-533]
	_ = x[_WM_MOVING-534]
	_ = x[_WM_POWERBROADCAST-536]
	_ = x[_WM_DEVICECHANGE-537]
	_ = x[_WM_MDICREATE-544]
	_ = x[_WM_MDIDESTROY-545]
	_ = x[_WM_MDIACTIVATE-546]
	_ = x[_WM_MDIRESTORE-547]
	_ = x[_WM_MDINEXT-548]
	_ = x[_WM_MDIMAXIMIZE-549]
	_ = x[_WM_MDITILE-550]
	_ = x[_WM_MDICASCADE-551]
	_ = x[_WM_MDIICONARRANGE-552]
	_ = x[_WM_MDIGETACTIVE-553]
	_ = x[_WM_MDISETMENU-560]
	_ = x[_WM_ENTERSIZEMOVE-561]
	_ = x[_WM_EXITSIZEMOVE-562]
	_ = x[_WM_DROPFILES-563]
	_ = x[_WM_MDIREFRESHMENU-564]
	_ = x[_WM_IME_SETCONTEXT-641]
	_ = x[_WM_IME_NOTIFY-642]
	_ = x[_WM_IME_CONTROL-643]
	_ = x[_WM_IME_COMPOSITIONFULL-644]
	_ = x[_WM_IME_SELECT-645]
	_ = x[_WM_IME_CHAR-646]
	_ = x[_WM_IME_KEYDOWN-656]
	_ = x[_WM_IME_KEYUP-657]
	_ = x[_WM_MOUSEHOVER-673]
	_ = x[_WM_NCMOUSELEAVE-674]
	_ = x[_WM_MOUSELEAVE-675]
	_ = x[_WM_CUT-768]
	_ = x[_WM_COPY-769]
	_ = x[_WM_PASTE-770]
	_ = x[_WM_CLEAR-771]
	_ = x[_WM_UNDO-772]
	_ = x[_WM_RENDERFORMAT-773]
	_ = x[_WM_RENDERALLFORMATS-774]
	_ = x[_WM_DESTROYCLIPBOARD-775]
	_ = x[_WM_DRAWCLIPBOARD-776]
	_ = x[_WM_PAINTCLIPBOARD-777]
	_ = x[_WM_VSCROLLCLIPBOARD-778]
	_ = x[_WM_SIZECLIPBOARD-779]
	_ = x[_WM_ASKCBFORMATNAME-780]
	_ = x[_WM_CHANGECBCHAIN-781]
	_ = x[_WM_HSCROLLCLIPBOARD-782]
	_ = x[_WM_QUERYNEWPALETTE-783]
	_ = x[_WM_PALETTEISCHANGING-784]
	_ = x[_WM_PALETTECHANGED-785]
	_ = x[_WM_HOTKEY-786]
	_ = x[_WM_PRINT-791]
	_ = x[_WM_PRINTCLIENT-792]
	_ = x[_WM_HANDHELDFIRST-856]
	_ = x[_WM_HANDHELDLAST-863]
	_ = x[_WM_PENWINFIRST-896]
	_ = x[_WM_PENWINLAST-911]
	_ = x[_WM_COALESCE_FIRST-912]
	_ = x[_WM_COALESCE_LAST-927]
	_ = x[_WM_DDE_FIRST-992]
	_ = x[_WM_DDE_INITIATE-992]
	_ = x[_WM_DDE_TERMINATE-993]
	_ = x[_WM_DDE_ADVISE-994]
	_ = x[_WM_DDE_UNADVISE-995]
	_ = x[_WM_DDE_ACK-996]
	_ = x[_WM_DDE_DATA-997]
	_ = x[_WM_DDE_REQUEST-998]
	_ = x[_WM_DDE_POKE-999]
	_ = x[_WM_DDE_EXECUTE-1000]
	_ = x[_WM_DDE_LAST-1000]
	_ = x[_WM_USER-1024]
	_ = x[_WM_APP-32768]
}

const __wm_name = "_WM_NULL_WM_CREATE_WM_DESTROY_WM_MOVE_WM_SIZE_WM_ACTIVATE_WM_SETFOCUS_WM_KILLFOCUS_WM_ENABLE_WM_SETREDRAW_WM_SETTEXT_WM_GETTEXT_WM_GETTEXTLENGTH_WM_PAINT_WM_CLOSE_WM_QUERYENDSESSION_WM_QUIT_WM_QUERYOPEN_WM_ERASEBKGND_WM_SYSCOLORCHANGE_WM_ENDSESSION_WM_SYSTEMERROR_WM_SHOWWINDOW_WM_CTLCOLOR_WM_WININICHANGE_WM_DEVMODECHANGE_WM_ACTIVATEAPP_WM_FONTCHANGE_WM_TIMECHANGE_WM_CANCELMODE_WM_SETCURSOR_WM_MOUSEACTIVATE_WM_CHILDACTIVATE_WM_QUEUESYNC_WM_GETMINMAXINFO_WM_PAINTICON_WM_ICONERASEBKGND_WM_NEXTDLGCTL_WM_SPOOLERSTATUS_WM_DRAWITEM_WM_MEASUREITEM_WM_DELETEITEM_WM_VKEYTOITEM_WM_CHARTOITEM_WM_SETFONT_WM_GETFONT_WM_SETHOTKEY_WM_GETHOTKEY_WM_QUERYDRAGICON_WM_COMPAREITEM_WM_COMPACTING_WM_WINDOWPOSCHANGING_WM_WINDOWPOSCHANGED_WM_POWER_WM_COPYDATA_WM_CANCELJOURNAL_WM_NOTIFY_WM_INPUTLANGCHANGEREQUEST_WM_INPUTLANGCHANGE_WM_TCARD_WM_HELP_WM_USERCHANGED_WM_NOTIFYFORMAT_WM_CONTEXTMENU_WM_STYLECHANGING_WM_STYLECHANGED_WM_DISPLAYCHANGE_WM_GETICON_WM_SETICON_WM_NCCREATE_WM_NCDESTROY_WM_NCCALCSIZE_WM_NCHITTEST_WM_NCPAINT_WM_NCACTIVATE_WM_GETDLGCODE_WM_NCMOUSEMOVE_WM_NCLBUTTONDOWN_WM_NCLBUTTONUP_WM_NCLBUTTONDBLCLK_WM_NCRBUTTONDOWN_WM_NCRBUTTONUP_WM_NCRBUTTONDBLCLK_WM_NCMBUTTONDOWN_WM_NCMBUTTONUP_WM_NCMBUTTONDBLCLK_WM_KEYDOWN_WM_KEYUP_WM_CHAR_WM_DEADCHAR_WM_SYSKEYDOWN_WM_SYSKEYUP_WM_SYSCHAR_WM_SYSDEADCHAR_WM_KEYLAST_WM_IME_STARTCOMPOSITION_WM_IME_ENDCOMPOSITION_WM_IME_COMPOSITION_WM_INITDIALOG_WM_COMMAND_WM_SYSCOMMAND_WM_TIMER_WM_HSCROLL_WM_VSCROLL_WM_INITMENU_WM_INITMENUPOPUP_WM_MENUSELECT_WM_MENUCHAR_WM_ENTERIDLE_WM_CTLCOLORMSGBOX_WM_CTLCOLOREDIT_WM_CTLCOLORLISTBOX_WM_CTLCOLORBTN_WM_CTLCOLORDLG_WM_CTLCOLORSCROLLBAR_WM_CTLCOLORSTATIC_WM_MOUSEMOVE_WM_LBUTTONDOWN_WM_LBUTTONUP_WM_LBUTTONDBLCLK_WM_RBUTTONDOWN_WM_RBUTTONUP_WM_RBUTTONDBLCLK_WM_MBUTTONDOWN_WM_MBUTTONUP_WM_MBUTTONDBLCLK_WM_MOUSEWHEEL_WM_MOUSEHWHEEL_WM_PARENTNOTIFY_WM_ENTERMENULOOP_WM_EXITMENULOOP_WM_NEXTMENU_WM_SIZING_WM_CAPTURECHANGED_WM_MOVING_WM_POWERBROADCAST_WM_DEVICECHANGE_WM_MDICREATE_WM_MDIDESTROY_WM_MDIACTIVATE_WM_MDIRESTORE_WM_MDINEXT_WM_MDIMAXIMIZE_WM_MDITILE_WM_MDICASCADE_WM_MDIICONARRANGE_WM_MDIGETACTIVE_WM_MDISETMENU_WM_ENTERSIZEMOVE_WM_EXITSIZEMOVE_WM_DROPFILES_WM_MDIREFRESHMENU_WM_IME_SETCONTEXT_WM_IME_NOTIFY_WM_IME_CONTROL_WM_IME_COMPOSITIONFULL_WM_IME_SELECT_WM_IME_CHAR_WM_IME_KEYDOWN_WM_IME_KEYUP_WM_MOUSEHOVER_WM_NCMOUSELEAVE_WM_MOUSELEAVE_WM_CUT_WM_COPY_WM_PASTE_WM_CLEAR_WM_UNDO_WM_RENDERFORMAT_WM_RENDERALLFORMATS_WM_DESTROYCLIPBOARD_WM_DRAWCLIPBOARD_WM_PAINTCLIPBOARD_WM_VSCROLLCLIPBOARD_WM_SIZECLIPBOARD_WM_ASKCBFORMATNAME_WM_CHANGECBCHAIN_WM_HSCROLLCLIPBOARD_WM_QUERYNEWPALETTE_WM_PALETTEISCHANGING_WM_PALETTECHANGED_WM_HOTKEY_WM_PRINT_WM_PRINTCLIENT_WM_HANDHELDFIRST_WM_HANDHELDLAST_WM_PENWINFIRST_WM_PENWINLAST_WM_COALESCE_FIRST_WM_COALESCE_LAST_WM_DDE_FIRST_WM_DDE_TERMINATE_WM_DDE_ADVISE_WM_DDE_UNADVISE_WM_DDE_ACK_WM_DDE_DATA_WM_DDE_REQUEST_WM_DDE_POKE_WM_DDE_EXECUTE_WM_USER_WM_APP"

var __wm_map = map[_wm]string{
	0:     __wm_name[0:8],
	1:     __wm_name[8:18],
	2:     __wm_name[18:29],
	3:     __wm_name[29:37],
	5:     __wm_name[37:45],
	6:     __wm_name[45:57],
	7:     __wm_name[57:69],
	8:     __wm_name[69:82],
	10:    __wm_name[82:92],
	11:    __wm_name[92:105],
	12:    __wm_name[105:116],
	13:    __wm_name[116:127],
	14:    __wm_name[127:144],
	15:    __wm_name[144:153],
	16:    __wm_name[153:162],
	17:    __wm_name[162:181],
	18:    __wm_name[181:189],
	19:    __wm_name[189:202],
	20:    __wm_name[202:216],
	21:    __wm_name[216:234],
	22:    __wm_name[234:248],
	23:    __wm_name[248:263],
	24:    __wm_name[263:277],
	25:    __wm_name[277:289],
	26:    __wm_name[289:305],
	27:    __wm_name[305:322],
	28:    __wm_name[322:337],
	29:    __wm_name[337:351],
	30:    __wm_name[351:365],
	31:    __wm_name[365:379],
	32:    __wm_name[379:392],
	33:    __wm_name[392:409],
	34:    __wm_name[409:426],
	35:    __wm_name[426:439],
	36:    __wm_name[439:456],
	38:    __wm_name[456:469],
	39:    __wm_name[469:487],
	40:    __wm_name[487:501],
	42:    __wm_name[501:518],
	43:    __wm_name[518:530],
	44:    __wm_name[530:545],
	45:    __wm_name[545:559],
	46:    __wm_name[559:573],
	47:    __wm_name[573:587],
	48:    __wm_name[587:598],
	49:    __wm_name[598:609],
	50:    __wm_name[609:622],
	51:    __wm_name[622:635],
	55:    __wm_name[635:652],
	57:    __wm_name[652:667],
	65:    __wm_name[667:681],
	70:    __wm_name[681:702],
	71:    __wm_name[702:722],
	72:    __wm_name[722:731],
	74:    __wm_name[731:743],
	75:    __wm_name[743:760],
	78:    __wm_name[760:770],
	80:    __wm_name[770:796],
	81:    __wm_name[796:815],
	82:    __wm_name[815:824],
	83:    __wm_name[824:832],
	84:    __wm_name[832:847],
	85:    __wm_name[847:863],
	123:   __wm_name[863:878],
	124:   __wm_name[878:895],
	125:   __wm_name[895:911],
	126:   __wm_name[911:928],
	127:   __wm_name[928:939],
	128:   __wm_name[939:950],
	129:   __wm_name[950:962],
	130:   __wm_name[962:975],
	131:   __wm_name[975:989],
	132:   __wm_name[989:1002],
	133:   __wm_name[1002:1013],
	134:   __wm_name[1013:1027],
	135:   __wm_name[1027:1041],
	160:   __wm_name[1041:1056],
	161:   __wm_name[1056:1073],
	162:   __wm_name[1073:1088],
	163:   __wm_name[1088:1107],
	164:   __wm_name[1107:1124],
	165:   __wm_name[1124:1139],
	166:   __wm_name[1139:1158],
	167:   __wm_name[1158:1175],
	168:   __wm_name[1175:1190],
	169:   __wm_name[1190:1209],
	256:   __wm_name[1209:1220],
	257:   __wm_name[1220:1229],
	258:   __wm_name[1229:1237],
	259:   __wm_name[1237:1249],
	260:   __wm_name[1249:1263],
	261:   __wm_name[1263:1275],
	262:   __wm_name[1275:1286],
	263:   __wm_name[1286:1301],
	264:   __wm_name[1301:1312],
	269:   __wm_name[1312:1336],
	270:   __wm_name[1336:1358],
	271:   __wm_name[1358:1377],
	272:   __wm_name[1377:1391],
	273:   __wm_name[1391:1402],
	274:   __wm_name[1402:1416],
	275:   __wm_name[1416:1425],
	276:   __wm_name[1425:1436],
	277:   __wm_name[1436:1447],
	278:   __wm_name[1447:1459],
	279:   __wm_name[1459:1476],
	287:   __wm_name[1476:1490],
	288:   __wm_name[1490:1502],
	289:   __wm_name[1502:1515],
	306:   __wm_name[1515:1533],
	307:   __wm_name[1533:1549],
	308:   __wm_name[1549:1568],
	309:   __wm_name[1568:1583],
	310:   __wm_name[1583:1598],
	311:   __wm_name[1598:1619],
	312:   __wm_name[1619:1637],
	512:   __wm_name[1637:1650],
	513:   __wm_name[1650:1665],
	514:   __wm_name[1665:1678],
	515:   __wm_name[1678:1695],
	516:   __wm_name[1695:1710],
	517:   __wm_name[1710:1723],
	518:   __wm_name[1723:1740],
	519:   __wm_name[1740:1755],
	520:   __wm_name[1755:1768],
	521:   __wm_name[1768:1785],
	522:   __wm_name[1785:1799],
	526:   __wm_name[1799:1814],
	528:   __wm_name[1814:1830],
	529:   __wm_name[1830:1847],
	530:   __wm_name[1847:1863],
	531:   __wm_name[1863:1875],
	532:   __wm_name[1875:1885],
	533:   __wm_name[1885:1903],
	534:   __wm_name[1903:1913],
	536:   __wm_name[1913:1931],
	537:   __wm_name[1931:1947],
	544:   __wm_name[1947:1960],
	545:   __wm_name[1960:1974],
	546:   __wm_name[1974:1989],
	547:   __wm_name[1989:2003],
	548:   __wm_name[2003:2014],
	549:   __wm_name[2014:2029],
	550:   __wm_name[2029:2040],
	551:   __wm_name[2040:2054],
	552:   __wm_name[2054:2072],
	553:   __wm_name[2072:2088],
	560:   __wm_name[2088:2102],
	561:   __wm_name[2102:2119],
	562:   __wm_name[2119:2135],
	563:   __wm_name[2135:2148],
	564:   __wm_name[2148:2166],
	641:   __wm_name[2166:2184],
	642:   __wm_name[2184:2198],
	643:   __wm_name[2198:2213],
	644:   __wm_name[2213:2236],
	645:   __wm_name[2236:2250],
	646:   __wm_name[2250:2262],
	656:   __wm_name[2262:2277],
	657:   __wm_name[2277:2290],
	673:   __wm_name[2290:2304],
	674:   __wm_name[2304:2320],
	675:   __wm_name[2320:2334],
	768:   __wm_name[2334:2341],
	769:   __wm_name[2341:2349],
	770:   __wm_name[2349:2358],
	771:   __wm_name[2358:2367],
	772:   __wm_name[2367:2375],
	773:   __wm_name[2375:2391],
	774:   __wm_name[2391:2411],
	775:   __wm_name[2411:2431],
	776:   __wm_name[2431:2448],
	777:   __wm_name[2448:2466],
	778:   __wm_name[2466:2486],
	779:   __wm_name[2486:2503],
	780:   __wm_name[2503:2522],
	781:   __wm_name[2522:2539],
	782:   __wm_name[2539:2559],
	783:   __wm_name[2559:2578],
	784:   __wm_name[2578:2599],
	785:   __wm_name[2599:2617],
	786:   __wm_name[2617:2627],
	791:   __wm_name[2627:2636],
	792:   __wm_name[2636:2651],
	856:   __wm_name[2651:2668],
	863:   __wm_name[2668:2684],
	896:   __wm_name[2684:2699],
	911:   __wm_name[2699:2713],
	912:   __wm_name[2713:2731],
	927:   __wm_name[2731:2748],
	992:   __wm_name[2748:2761],
	993:   __wm_name[2761:2778],
	994:   __wm_name[2778:2792],
	995:   __wm_name[2792:2808],
	996:   __wm_name[2808:2819],
	997:   __wm_name[2819:2831],
	998:   __wm_name[2831:2846],
	999:   __wm_name[2846:2858],
	1000:  __wm_name[2858:2873],
	1024:  __wm_name[2873:2881],
	32768: __wm_name[2881:2888],
}

func (i _wm) String() string {
	if str, ok := __wm_map[i]; ok {
		return str
	}
	return "_wm(" + strconv.FormatInt(int64(i), 10) + ")"
}