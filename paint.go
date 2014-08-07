package main

import (
	"github.com/nsf/termbox-go"
)

var repaint = make(chan struct{}, 1)

// Repaints queues a repaint if there is not one already queued.
func Repaint() {
	select {
	case repaint <- struct{}{}:
	default:
	}
}

func paint() {
	InitializeUI()

	w, h := termbox.Size()

	UI.Paint(0, 0, 0, 0, w, h)

	termbox.Flush()
}

func InitializeUI() {
	if UI != nil {
		return
	}

	UI = &RootPanel{
		Child: &BorderPanel{
			Title: []rune("Suudsu"),
			Fg:    termbox.ColorBlack,
			Bg:    termbox.ColorWhite,
		},
	}
}
