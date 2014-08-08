package main

import (
	"github.com/nsf/termbox-go"
)

type Panel interface {
	Paint(offsetX, offsetY, left, top, width, height int)

	HandleKey(key termbox.Key, mod termbox.Modifier) bool
	HandleRune(r rune, mod termbox.Modifier) bool
	HandleMouse(x, y int) bool
}

var UI Panel

type RootPanel struct {
	Child Panel
}

var _ Panel = (*RootPanel)(nil)

func (p *RootPanel) Paint(offsetX, offsetY, left, top, width, height int) {
	if width <= 0 || height <= 0 {
		return
	}

	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)

	if p.Child != nil {
		p.Child.Paint(offsetX, offsetY, left, top, width, height)
	}
}

func (p *RootPanel) HandleKey(key termbox.Key, mod termbox.Modifier) bool {
	if p.Child != nil {
		return p.Child.HandleKey(key, mod)
	}
	return false
}

func (p *RootPanel) HandleRune(r rune, mod termbox.Modifier) bool {
	if p.Child != nil {
		return p.Child.HandleRune(r, mod)
	}
	return false
}

func (p *RootPanel) HandleMouse(x, y int) bool {
	if p.Child != nil {
		return p.Child.HandleMouse(x, y)
	}
	return false
}

type BorderPanel struct {
	Title  []rune
	Fg, Bg termbox.Attribute
	Child  Panel
}

var _ Panel = (*BorderPanel)(nil)

func (p *BorderPanel) Paint(offsetX, offsetY, left, top, width, height int) {
	if width <= 0 || height <= 0 {
		return
	}

	titleOffset := (width - len(p.Title)) / 2

	for i := 0; i < height; i++ {
		termbox.SetCell(left, top+i, '│', p.Fg, p.Bg)
		termbox.SetCell(left+width-1, top+i, '│', p.Fg, p.Bg)
	}
	for i := 0; i < width; i++ {
		r := '─'
		bottom := '─'
		if i == 0 {
			r = '┌'
			bottom = '└'
		} else if i == width-1 {
			r = '┐'
			bottom = '┘'
		}
		if i >= titleOffset && i < titleOffset+len(p.Title) {
			r = p.Title[i-titleOffset]
		}
		termbox.SetCell(left+i, top, r, p.Fg, p.Bg)
		termbox.SetCell(left+i, top+height-1, bottom, p.Fg, p.Bg)
	}
	if p.Child != nil {
		p.Child.Paint(offsetX, offsetY, left+1, top+1, width-2, height-2)
	}
}

func (p *BorderPanel) HandleKey(key termbox.Key, mod termbox.Modifier) bool {
	if p.Child != nil {
		return p.Child.HandleKey(key, mod)
	}
	return false
}

func (p *BorderPanel) HandleRune(r rune, mod termbox.Modifier) bool {
	if p.Child != nil {
		return p.Child.HandleRune(r, mod)
	}
	return false
}

func (p *BorderPanel) HandleMouse(x, y int) bool {
	if p.Child != nil {
		return p.Child.HandleMouse(x-1, y-1)
	}
	return false
}
