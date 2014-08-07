package main

import (
	"github.com/nsf/termbox-go"
)

var exit = make(chan struct{})

func HandleKey(key termbox.Key, mod termbox.Modifier) {
	if key == termbox.KeyEsc {
		close(exit)
	}
}

func HandleRune(r rune, mod termbox.Modifier) {

}

func HandleMouse(x, y int) {

}
