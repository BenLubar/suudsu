package main

import (
	"github.com/nsf/termbox-go"
	"github.com/robertkrimen/otto"
	"log"
	"reflect"
)

func ScriptGlobals(o *otto.Otto) {
	t, err := o.Object(`termbox = {}`)
	if err != nil {
		log.Panicln("ScriptGlobals:", err)
	}
	t.Set("KeyF1", int(termbox.KeyF1))
	t.Set("KeyF2", int(termbox.KeyF2))
	t.Set("KeyF3", int(termbox.KeyF3))
	t.Set("KeyF4", int(termbox.KeyF4))
	t.Set("KeyF5", int(termbox.KeyF5))
	t.Set("KeyF6", int(termbox.KeyF6))
	t.Set("KeyF7", int(termbox.KeyF7))
	t.Set("KeyF8", int(termbox.KeyF8))
	t.Set("KeyF9", int(termbox.KeyF9))
	t.Set("KeyF10", int(termbox.KeyF10))
	t.Set("KeyF11", int(termbox.KeyF11))
	t.Set("KeyF12", int(termbox.KeyF12))
	t.Set("KeyInsert", int(termbox.KeyInsert))
	t.Set("KeyDelete", int(termbox.KeyDelete))
	t.Set("KeyHome", int(termbox.KeyHome))
	t.Set("KeyEnd", int(termbox.KeyEnd))
	t.Set("KeyPgup", int(termbox.KeyPgup))
	t.Set("KeyPgdn", int(termbox.KeyPgdn))
	t.Set("KeyArrowUp", int(termbox.KeyArrowUp))
	t.Set("KeyArrowDown", int(termbox.KeyArrowDown))
	t.Set("KeyArrowLeft", int(termbox.KeyArrowLeft))
	t.Set("KeyArrowRight", int(termbox.KeyArrowRight))
	t.Set("MouseLeft", int(termbox.MouseLeft))
	t.Set("MouseMiddle", int(termbox.MouseMiddle))
	t.Set("MouseRight", int(termbox.MouseRight))
	t.Set("KeyCtrlTilde", int(termbox.KeyCtrlTilde))
	t.Set("KeyCtrl2", int(termbox.KeyCtrl2))
	t.Set("KeyCtrlSpace", int(termbox.KeyCtrlSpace))
	t.Set("KeyCtrlA", int(termbox.KeyCtrlA))
	t.Set("KeyCtrlB", int(termbox.KeyCtrlB))
	t.Set("KeyCtrlC", int(termbox.KeyCtrlC))
	t.Set("KeyCtrlD", int(termbox.KeyCtrlD))
	t.Set("KeyCtrlE", int(termbox.KeyCtrlE))
	t.Set("KeyCtrlF", int(termbox.KeyCtrlF))
	t.Set("KeyCtrlG", int(termbox.KeyCtrlG))
	t.Set("KeyBackspace", int(termbox.KeyBackspace))
	t.Set("KeyCtrlH", int(termbox.KeyCtrlH))
	t.Set("KeyTab", int(termbox.KeyTab))
	t.Set("KeyCtrlI", int(termbox.KeyCtrlI))
	t.Set("KeyCtrlJ", int(termbox.KeyCtrlJ))
	t.Set("KeyCtrlK", int(termbox.KeyCtrlK))
	t.Set("KeyCtrlL", int(termbox.KeyCtrlL))
	t.Set("KeyEnter", int(termbox.KeyEnter))
	t.Set("KeyCtrlM", int(termbox.KeyCtrlM))
	t.Set("KeyCtrlN", int(termbox.KeyCtrlN))
	t.Set("KeyCtrlO", int(termbox.KeyCtrlO))
	t.Set("KeyCtrlP", int(termbox.KeyCtrlP))
	t.Set("KeyCtrlQ", int(termbox.KeyCtrlQ))
	t.Set("KeyCtrlR", int(termbox.KeyCtrlR))
	t.Set("KeyCtrlS", int(termbox.KeyCtrlS))
	t.Set("KeyCtrlT", int(termbox.KeyCtrlT))
	t.Set("KeyCtrlU", int(termbox.KeyCtrlU))
	t.Set("KeyCtrlV", int(termbox.KeyCtrlV))
	t.Set("KeyCtrlW", int(termbox.KeyCtrlW))
	t.Set("KeyCtrlX", int(termbox.KeyCtrlX))
	t.Set("KeyCtrlY", int(termbox.KeyCtrlY))
	t.Set("KeyCtrlZ", int(termbox.KeyCtrlZ))
	t.Set("KeyEsc", int(termbox.KeyEsc))
	t.Set("KeyCtrlLsqBracket", int(termbox.KeyCtrlLsqBracket))
	t.Set("KeyCtrl3", int(termbox.KeyCtrl3))
	t.Set("KeyCtrl4", int(termbox.KeyCtrl4))
	t.Set("KeyCtrlBackslash", int(termbox.KeyCtrlBackslash))
	t.Set("KeyCtrl5", int(termbox.KeyCtrl5))
	t.Set("KeyCtrlRsqBracket", int(termbox.KeyCtrlRsqBracket))
	t.Set("KeyCtrl6", int(termbox.KeyCtrl6))
	t.Set("KeyCtrl7", int(termbox.KeyCtrl7))
	t.Set("KeyCtrlSlash", int(termbox.KeyCtrlSlash))
	t.Set("KeyCtrlUnderscore", int(termbox.KeyCtrlUnderscore))
	t.Set("KeySpace", int(termbox.KeySpace))
	t.Set("KeyBackspace2", int(termbox.KeyBackspace2))
	t.Set("KeyCtrl8", int(termbox.KeyCtrl8))
	t.Set("ModAlt", int(termbox.ModAlt))
	t.Set("ColorDefault", int(termbox.ColorDefault))
	t.Set("ColorBlack", int(termbox.ColorBlack))
	t.Set("ColorRed", int(termbox.ColorRed))
	t.Set("ColorGreen", int(termbox.ColorGreen))
	t.Set("ColorYellow", int(termbox.ColorYellow))
	t.Set("ColorBlue", int(termbox.ColorBlue))
	t.Set("ColorMagenta", int(termbox.ColorMagenta))
	t.Set("ColorCyan", int(termbox.ColorCyan))
	t.Set("ColorWhite", int(termbox.ColorWhite))
	t.Set("AttrBold", int(termbox.AttrBold))
	t.Set("AttrUnderline", int(termbox.AttrUnderline))
	t.Set("AttrReverse", int(termbox.AttrReverse))
	t.Set("SetCell", func(f otto.FunctionCall) otto.Value {
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
	})

	o.Set("UI", func(f otto.FunctionCall) otto.Value {
		InitializeUI()

		p, _ := f.Otto.Object(`{}`)
		initScriptPanel(p, UI)
		return p.Value()
	})

	o.Set("BorderPanel", func(f otto.FunctionCall) otto.Value {
		initScriptPanel(f.This.Object(), &BorderPanel{})
		return f.This
	})
}

var float64Type = reflect.TypeOf(float64(0))
var stringType = reflect.TypeOf(string(""))
var boolType = reflect.TypeOf(bool(false))
var panelType reflect.Type

func init() {
	var panel Panel
	panelType = reflect.TypeOf(&panel).Elem()
}

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
		key, err := f.Argument(0).ToInteger()
		if err != nil {
			panic(err)
		}
		mod, err := f.Argument(1).ToInteger()
		if err != nil {
			panic(err)
		}
		v, err := otto.ToValue(p.HandleKey(termbox.Key(key), termbox.Modifier(mod)))
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
		mod, err := f.Argument(1).ToInteger()
		if err != nil {
			panic(err)
		}
		r := []rune(s)
		if len(r) != 1 {
			panic("HandleRune: invalid value for rune")
		}
		v, err := otto.ToValue(p.HandleRune(r[0], termbox.Modifier(mod)))
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

	vp := reflect.ValueOf(p)
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

	if fn.IsFunction() {
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

func (p *ScriptedPanel) HandleKey(key termbox.Key, mod termbox.Modifier) bool {
	v := p.fn("HandleKey", int(key), int(mod))
	b, err := v.ToBoolean()
	if err != nil {
		log.Println("ScriptedPanel:", "HandleKey", p.Script, err)
		return false
	}
	return b
}

func (p *ScriptedPanel) HandleRune(r rune, mod termbox.Modifier) bool {
	v := p.fn("HandleRune", string(r), int(mod))
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
