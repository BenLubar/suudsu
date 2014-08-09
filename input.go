package main

import (
	"github.com/nsf/termbox-go"
	"github.com/robertkrimen/otto"
	"log"
	"os"
)

func HandleKey(key termbox.Key, mod termbox.Modifier) {
	InitializeUI()

	handled := false
	for _, bind := range binds.Keys[keyBind{K: key, M: mod}] {
		handled = UI.HandleKey(bind) || handled
	}

	if !handled && key == termbox.KeyEsc && mod == 0 {
		select {
		case exit <- struct{}{}:
		default:
		}
	}
}

func HandleRune(r rune, mod termbox.Modifier) {
	InitializeUI()

	handled := false
	if mod == 0 {
		handled = UI.HandleRune(r) || handled
	}
	for _, bind := range binds.Runes[runeBind{R: r, M: mod}] {
		handled = UI.HandleKey(bind) || handled
	}

	_ = handled
}

func HandleMouse(x, y int) {
	InitializeUI()

	UI.HandleMouse(x, y)
}

type keyBind struct {
	K termbox.Key
	M termbox.Modifier
}

type runeBind struct {
	R rune
	M termbox.Modifier
}

var binds struct {
	Keys  map[keyBind][]string
	Runes map[runeBind][]string
}

func init() {
	LoadBindings()
}

func LoadBindings() {
	binds.Keys = make(map[keyBind][]string)
	binds.Runes = make(map[runeBind][]string)

	o := otto.New()
	ScriptGlobals(o)
	o.Set("BindKey", func(f otto.FunctionCall) otto.Value {
		k, err := f.Argument(0).ToInteger()
		if err != nil {
			panic(err)
		}
		m, err := f.Argument(1).ToInteger()
		if err != nil {
			panic(err)
		}
		b, err := f.Argument(2).ToString()
		if err != nil {
			panic(err)
		}

		kb := keyBind{
			K: termbox.Key(k),
			M: termbox.Modifier(m),
		}

		binds.Keys[kb] = append(binds.Keys[kb], b)
		return otto.UndefinedValue()
	})
	o.Set("BindRune", func(f otto.FunctionCall) otto.Value {
		s, err := f.Argument(0).ToString()
		if err != nil {
			panic(err)
		}
		m, err := f.Argument(1).ToInteger()
		if err != nil {
			panic(err)
		}
		b, err := f.Argument(2).ToString()
		if err != nil {
			panic(err)
		}

		r := []rune(s)
		if len(r) != 1 {
			panic("invalid rune")
		}

		rb := runeBind{
			R: r[0],
			M: termbox.Modifier(m),
		}

		binds.Runes[rb] = append(binds.Runes[rb], b)
		return otto.UndefinedValue()
	})

	f, err := os.Open("bindings.js")
	if err != nil {
		log.Println("LoadBindings:", err)
		return
	}
	defer f.Close()

	_, err = o.Run(f)
	if err != nil {
		log.Println("LoadBindings:", err)
	}
}
