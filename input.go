package main

import (
	"github.com/nsf/termbox-go"
)

var exit = make(chan struct{})

func HandleKey(key termbox.Key, mod termbox.Modifier) {
	if key == termbox.KeyEsc {
		close(exit)
	}

	InitializeUI()

	UI.HandleKey(key, mod)
}

func HandleRune(r rune, mod termbox.Modifier) {
	InitializeUI()

	UI.HandleRune(r, mod)
}

func HandleMouse(x, y int) {
	InitializeUI()

	UI.HandleMouse(x, y)
}
