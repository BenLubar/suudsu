package main

import (
	"github.com/nsf/termbox-go"
	"github.com/robertkrimen/otto"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

func ScriptGlobals(o *otto.Otto) {
	defer func() {
		if err := recover(); err != nil {
			log.Panicln("ScriptGlobals:", err)
		}
	}()
	checkError := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	t, err := o.Object(`termbox = {}`)
	checkError(err)
	checkError(t.Set("KeyF1", int(termbox.KeyF1)))
	checkError(t.Set("KeyF2", int(termbox.KeyF2)))
	checkError(t.Set("KeyF3", int(termbox.KeyF3)))
	checkError(t.Set("KeyF4", int(termbox.KeyF4)))
	checkError(t.Set("KeyF5", int(termbox.KeyF5)))
	checkError(t.Set("KeyF6", int(termbox.KeyF6)))
	checkError(t.Set("KeyF7", int(termbox.KeyF7)))
	checkError(t.Set("KeyF8", int(termbox.KeyF8)))
	checkError(t.Set("KeyF9", int(termbox.KeyF9)))
	checkError(t.Set("KeyF10", int(termbox.KeyF10)))
	checkError(t.Set("KeyF11", int(termbox.KeyF11)))
	checkError(t.Set("KeyF12", int(termbox.KeyF12)))
	checkError(t.Set("KeyInsert", int(termbox.KeyInsert)))
	checkError(t.Set("KeyDelete", int(termbox.KeyDelete)))
	checkError(t.Set("KeyHome", int(termbox.KeyHome)))
	checkError(t.Set("KeyEnd", int(termbox.KeyEnd)))
	checkError(t.Set("KeyPgup", int(termbox.KeyPgup)))
	checkError(t.Set("KeyPgdn", int(termbox.KeyPgdn)))
	checkError(t.Set("KeyArrowUp", int(termbox.KeyArrowUp)))
	checkError(t.Set("KeyArrowDown", int(termbox.KeyArrowDown)))
	checkError(t.Set("KeyArrowLeft", int(termbox.KeyArrowLeft)))
	checkError(t.Set("KeyArrowRight", int(termbox.KeyArrowRight)))
	checkError(t.Set("MouseLeft", int(termbox.MouseLeft)))
	checkError(t.Set("MouseMiddle", int(termbox.MouseMiddle)))
	checkError(t.Set("MouseRight", int(termbox.MouseRight)))
	checkError(t.Set("KeyCtrlTilde", int(termbox.KeyCtrlTilde)))
	checkError(t.Set("KeyCtrl2", int(termbox.KeyCtrl2)))
	checkError(t.Set("KeyCtrlSpace", int(termbox.KeyCtrlSpace)))
	checkError(t.Set("KeyCtrlA", int(termbox.KeyCtrlA)))
	checkError(t.Set("KeyCtrlB", int(termbox.KeyCtrlB)))
	checkError(t.Set("KeyCtrlC", int(termbox.KeyCtrlC)))
	checkError(t.Set("KeyCtrlD", int(termbox.KeyCtrlD)))
	checkError(t.Set("KeyCtrlE", int(termbox.KeyCtrlE)))
	checkError(t.Set("KeyCtrlF", int(termbox.KeyCtrlF)))
	checkError(t.Set("KeyCtrlG", int(termbox.KeyCtrlG)))
	checkError(t.Set("KeyBackspace", int(termbox.KeyBackspace)))
	checkError(t.Set("KeyCtrlH", int(termbox.KeyCtrlH)))
	checkError(t.Set("KeyTab", int(termbox.KeyTab)))
	checkError(t.Set("KeyCtrlI", int(termbox.KeyCtrlI)))
	checkError(t.Set("KeyCtrlJ", int(termbox.KeyCtrlJ)))
	checkError(t.Set("KeyCtrlK", int(termbox.KeyCtrlK)))
	checkError(t.Set("KeyCtrlL", int(termbox.KeyCtrlL)))
	checkError(t.Set("KeyEnter", int(termbox.KeyEnter)))
	checkError(t.Set("KeyCtrlM", int(termbox.KeyCtrlM)))
	checkError(t.Set("KeyCtrlN", int(termbox.KeyCtrlN)))
	checkError(t.Set("KeyCtrlO", int(termbox.KeyCtrlO)))
	checkError(t.Set("KeyCtrlP", int(termbox.KeyCtrlP)))
	checkError(t.Set("KeyCtrlQ", int(termbox.KeyCtrlQ)))
	checkError(t.Set("KeyCtrlR", int(termbox.KeyCtrlR)))
	checkError(t.Set("KeyCtrlS", int(termbox.KeyCtrlS)))
	checkError(t.Set("KeyCtrlT", int(termbox.KeyCtrlT)))
	checkError(t.Set("KeyCtrlU", int(termbox.KeyCtrlU)))
	checkError(t.Set("KeyCtrlV", int(termbox.KeyCtrlV)))
	checkError(t.Set("KeyCtrlW", int(termbox.KeyCtrlW)))
	checkError(t.Set("KeyCtrlX", int(termbox.KeyCtrlX)))
	checkError(t.Set("KeyCtrlY", int(termbox.KeyCtrlY)))
	checkError(t.Set("KeyCtrlZ", int(termbox.KeyCtrlZ)))
	checkError(t.Set("KeyEsc", int(termbox.KeyEsc)))
	checkError(t.Set("KeyCtrlLsqBracket", int(termbox.KeyCtrlLsqBracket)))
	checkError(t.Set("KeyCtrl3", int(termbox.KeyCtrl3)))
	checkError(t.Set("KeyCtrl4", int(termbox.KeyCtrl4)))
	checkError(t.Set("KeyCtrlBackslash", int(termbox.KeyCtrlBackslash)))
	checkError(t.Set("KeyCtrl5", int(termbox.KeyCtrl5)))
	checkError(t.Set("KeyCtrlRsqBracket", int(termbox.KeyCtrlRsqBracket)))
	checkError(t.Set("KeyCtrl6", int(termbox.KeyCtrl6)))
	checkError(t.Set("KeyCtrl7", int(termbox.KeyCtrl7)))
	checkError(t.Set("KeyCtrlSlash", int(termbox.KeyCtrlSlash)))
	checkError(t.Set("KeyCtrlUnderscore", int(termbox.KeyCtrlUnderscore)))
	checkError(t.Set("KeySpace", int(termbox.KeySpace)))
	checkError(t.Set("KeyBackspace2", int(termbox.KeyBackspace2)))
	checkError(t.Set("KeyCtrl8", int(termbox.KeyCtrl8)))
	checkError(t.Set("ModAlt", int(termbox.ModAlt)))
	checkError(t.Set("ColorDefault", int(termbox.ColorDefault)))
	checkError(t.Set("ColorBlack", int(termbox.ColorBlack)))
	checkError(t.Set("ColorRed", int(termbox.ColorRed)))
	checkError(t.Set("ColorGreen", int(termbox.ColorGreen)))
	checkError(t.Set("ColorYellow", int(termbox.ColorYellow)))
	checkError(t.Set("ColorBlue", int(termbox.ColorBlue)))
	checkError(t.Set("ColorMagenta", int(termbox.ColorMagenta)))
	checkError(t.Set("ColorCyan", int(termbox.ColorCyan)))
	checkError(t.Set("ColorWhite", int(termbox.ColorWhite)))
	checkError(t.Set("AttrBold", int(termbox.AttrBold)))
	checkError(t.Set("AttrUnderline", int(termbox.AttrUnderline)))
	checkError(t.Set("AttrReverse", int(termbox.AttrReverse)))
	checkError(t.Set("SetCell", func(f otto.FunctionCall) otto.Value {
		x, err := f.Argument(0).ToInteger()
		if err != nil {
			panic(err)
		}
		y, err := f.Argument(1).ToInteger()
		if err != nil {
			panic(err)
		}
		s, err := f.Argument(2).ToString()
		if err != nil {
			panic(err)
		}
		fg, err := f.Argument(3).ToInteger()
		if err != nil {
			panic(err)
		}
		bg, err := f.Argument(4).ToInteger()
		if err != nil {
			panic(err)
		}

		r := []rune(s)
		if len(r) != 1 {
			panic("SetCell: invalid value for rune")
		}
		termbox.SetCell(int(x), int(y), r[0], termbox.Attribute(fg), termbox.Attribute(bg))
		return otto.UndefinedValue()
	}))

	checkError(o.Set("UI", func(f otto.FunctionCall) otto.Value {
		InitializeUI()

		p, _ := f.Otto.Object(`({})`)
		initScriptPanel(p, UI)
		return p.Value()
	}))

	checkError(o.Set("BorderPanel", func(f otto.FunctionCall) otto.Value {
		initScriptPanel(f.This.Object(), &BorderPanel{})
		return f.This
	}))

	checkError(o.Set("Repaint", func(f otto.FunctionCall) otto.Value {
		Repaint()
		return otto.UndefinedValue()
	}))

	checkError(o.Set("require", func(f otto.FunctionCall) otto.Value {
		path := filepath.Clean(f.Argument(0).String() + ".js")

		r, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer r.Close()

		v, err := f.Otto.Run(r)
		if err != nil {
			panic(err)
		}

		return v
	}))

	checkError(o.Set("Exit", func(f otto.FunctionCall) otto.Value {
		select {
		case exit <- struct{}{}:
		default:
		}
		return otto.UndefinedValue()
	}))
}

var float64Type = reflect.TypeOf(float64(0))
var stringType = reflect.TypeOf(string(""))
var boolType = reflect.TypeOf(bool(false))
var panelType = func() reflect.Type {
	var panel Panel
	return reflect.TypeOf(&panel).Elem()
}()

func initScriptPanel(o *otto.Object, p Panel) {
	if sp, ok := p.(*ScriptedPanel); ok {
		for _, k := range sp.Script.Keys() {
			v, err := sp.Script.Get(k)
			if err == nil {
				err = o.Set(k, v)
			}
			if err != nil {
				log.Println("initScriptPanel:", err)
			}
		}
		return
	}

	err := o.Set("Paint", func(f otto.FunctionCall) otto.Value {
		offsetX, err := f.Argument(0).ToInteger()
		if err != nil {
			panic(err)
		}
		offsetY, err := f.Argument(1).ToInteger()
		if err != nil {
			panic(err)
		}
		left, err := f.Argument(2).ToInteger()
		if err != nil {
			panic(err)
		}
		top, err := f.Argument(3).ToInteger()
		if err != nil {
			panic(err)
		}
		width, err := f.Argument(4).ToInteger()
		if err != nil {
			panic(err)
		}
		height, err := f.Argument(5).ToInteger()
		if err != nil {
			panic(err)
		}
		p.Paint(int(offsetX), int(offsetY), int(left), int(top), int(width), int(height))
		return otto.UndefinedValue()
	})
	if err != nil {
		log.Println("initScriptPanel:", err)
	}
	err = o.Set("HandleKey", func(f otto.FunctionCall) otto.Value {
		key, err := f.Argument(0).ToString()
		if err != nil {
			panic(err)
		}
		v, err := otto.ToValue(p.HandleKey(key))
		if err != nil {
			panic(err)
		}
		return v
	})
	if err != nil {
		log.Println("initScriptPanel:", err)
	}
	err = o.Set("HandleRune", func(f otto.FunctionCall) otto.Value {
		s, err := f.Argument(0).ToString()
		if err != nil {
			panic(err)
		}
		r := []rune(s)
		if len(r) != 1 {
			panic("HandleRune: invalid value for rune")
		}
		v, err := otto.ToValue(p.HandleRune(r[0]))
		if err != nil {
			panic(err)
		}
		return v
	})
	if err != nil {
		log.Println("initScriptPanel:", err)
	}
	err = o.Set("HandleMouse", func(f otto.FunctionCall) otto.Value {
		x, err := f.Argument(0).ToInteger()
		if err != nil {
			panic(err)
		}
		y, err := f.Argument(1).ToInteger()
		if err != nil {
			panic(err)
		}
		v, err := otto.ToValue(p.HandleMouse(int(x), int(y)))
		if err != nil {
			panic(err)
		}
		return v
	})
	if err != nil {
		log.Println("initScriptPanel:", err)
	}

	vp := reflect.ValueOf(p).Elem()
	tp := vp.Type()
	for i := 0; i < tp.NumField(); i++ {
		ft := tp.Field(i)
		fv := vp.Field(i)
		if !fv.CanSet() {
			continue
		}
		if ft.Type.Implements(panelType) {
			err = o.Set("Get"+ft.Name, func(f otto.FunctionCall) otto.Value {
				panel, ok := fv.Interface().(Panel)
				if !ok {
					return otto.NullValue()
				}
				p, _ := f.Otto.Object(`{}`)
				initScriptPanel(p, panel)
				return p.Value()
			})
			if err != nil {
				log.Println("initScriptPanel:", err)
			}
			err = o.Set("Set"+ft.Name, func(f otto.FunctionCall) otto.Value {
				fv.Set(reflect.ValueOf(&ScriptedPanel{f.Argument(0).Object()}))
				return otto.UndefinedValue()
			})
			if err != nil {
				log.Println("initScriptPanel:", err)
			}
		} else if ft.Type.ConvertibleTo(float64Type) && float64Type.ConvertibleTo(ft.Type) {
			err = o.Set("Get"+ft.Name, func(f otto.FunctionCall) otto.Value {
				v, err := otto.ToValue(fv.Convert(float64Type).Float())
				if err != nil {
					panic(err)
				}
				return v
			})
			if err != nil {
				log.Println("initScriptPanel:", err)
			}
			err = o.Set("Set"+ft.Name, func(f otto.FunctionCall) otto.Value {
				v, err := f.Argument(0).ToFloat()
				if err != nil {
					panic(err)
				}
				fv.Set(reflect.ValueOf(v).Convert(ft.Type))
				return otto.UndefinedValue()
			})
			if err != nil {
				log.Println("initScriptPanel:", err)
			}
		} else if ft.Type.ConvertibleTo(stringType) && stringType.ConvertibleTo(ft.Type) {
			err = o.Set("Get"+ft.Name, func(f otto.FunctionCall) otto.Value {
				v, err := otto.ToValue(fv.Convert(stringType).String())
				if err != nil {
					panic(err)
				}
				return v
			})
			if err != nil {
				log.Println("initScriptPanel:", err)
			}
			err = o.Set("Set"+ft.Name, func(f otto.FunctionCall) otto.Value {
				s, err := f.Argument(0).ToString()
				if err != nil {
					panic(err)
				}
				fv.Set(reflect.ValueOf(s).Convert(ft.Type))
				return otto.UndefinedValue()
			})
			if err != nil {
				log.Println("initScriptPanel:", err)
			}
		} else if ft.Type.ConvertibleTo(boolType) && boolType.ConvertibleTo(ft.Type) {
			err = o.Set("Is"+ft.Name, func(f otto.FunctionCall) otto.Value {
				v, err := otto.ToValue(fv.Convert(boolType).Bool())
				if err != nil {
					panic(err)
				}
				return v
			})
			if err != nil {
				log.Println("initScriptPanel:", err)
			}
			err = o.Set("Set"+ft.Name, func(f otto.FunctionCall) otto.Value {
				b, err := f.Argument(0).ToBoolean()
				if err != nil {
					panic(err)
				}
				fv.Set(reflect.ValueOf(b).Convert(ft.Type))
				return otto.UndefinedValue()
			})
			if err != nil {
				log.Println("initScriptPanel:", err)
			}
		}
	}
}

type ScriptedPanel struct {
	Script *otto.Object
}

var _ Panel = (*ScriptedPanel)(nil)

func (p *ScriptedPanel) fn(name string, args ...interface{}) otto.Value {
	if p.Script == nil {
		log.Println("ScriptedPanel:", name, "missing script")
		return otto.UndefinedValue()
	}

	fn, err := p.Script.Get(name)
	if err != nil {
		log.Println("ScriptedPanel:", name, p.Script, err)
		return otto.UndefinedValue()
	}

	if !fn.IsFunction() {
		log.Println("ScriptedPanel:", name, p.Script, fn, "is not a function")
		return otto.UndefinedValue()
	}

	v, err := fn.Call(p.Script.Value(), args...)
	if err != nil {
		log.Println("ScriptedPanel:", name, p.Script, fn, err)
		return otto.UndefinedValue()
	}

	return v
}

func (p *ScriptedPanel) Paint(offsetX, offsetY, left, top, width, height int) {
	p.fn("Paint", offsetX, offsetY, left, top, width, height)
}

func (p *ScriptedPanel) HandleKey(key string) bool {
	v := p.fn("HandleKey", key)
	b, err := v.ToBoolean()
	if err != nil {
		log.Println("ScriptedPanel:", "HandleKey", p.Script, err)
		return false
	}
	return b
}

func (p *ScriptedPanel) HandleRune(r rune) bool {
	v := p.fn("HandleRune", string(r))
	b, err := v.ToBoolean()
	if err != nil {
		log.Println("ScriptedPanel:", "HandleRune", p.Script, err)
		return false
	}
	return b
}

func (p *ScriptedPanel) HandleMouse(x, y int) bool {
	v := p.fn("HandleMouse", x, y)
	b, err := v.ToBoolean()
	if err != nil {
		log.Println("ScriptedPanel:", "HandleMouse", p.Script, err)
		return false
	}
	return b
}
