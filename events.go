package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
)

func main() {
	if err := termbox.Init(); err != nil {
		fmt.Fprintln(os.Stderr, "Initialization error:", err)
		return
	}
	defer termbox.Close()

	events := make(chan termbox.Event)
	go pumpEvents(events)

	Repaint()

	for {
		select {
		case event := <-events:
			switch event.Type {
			case termbox.EventError:
				panic(event.Err)

			case termbox.EventResize:
				Repaint()

			case termbox.EventKey:
				if event.Ch == 0 {
					HandleKey(event.Key, event.Mod)
				} else {
					HandleRune(event.Ch, event.Mod)
				}

			case termbox.EventMouse:
				HandleMouse(event.MouseX, event.MouseY)
			}

		case <-repaint:
			paint()

		case <-exit:
			return
		}
	}
}

func pumpEvents(ch chan<- termbox.Event) {
	for {
		ch <- termbox.PollEvent()
	}
}
