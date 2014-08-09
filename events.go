package main

import (
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"syscall"
)

var exit = make(chan struct{}, 1)

func main() {
	redirectStderr("error.log")

	if err := termbox.Init(); err != nil {
		log.Println("Initialization error:", err)
		return
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

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

func redirectStderr(name string) {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("could not open error log for writing:", err)
	}
	defer f.Close()

	err = syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Fatalln("could not open error log for writing:", err)
	}
}
