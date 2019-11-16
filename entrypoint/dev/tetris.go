// +build dev

package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"tetris/pkg/game"

	"github.com/nsf/termbox-go"
)

func main() {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	tetris := game.DevRestore()
	defer tetris.DevDump()

	go func() {
		userInput := tetris.Actions()

		for {
			event := termbox.PollEvent()

			switch event.Key {
			case termbox.KeyEsc:
				tetris.Close()
			case termbox.KeySpace:
				userInput <- game.ActionPause
			case termbox.KeyArrowUp:
				userInput <- game.ActionUp
			case termbox.KeyArrowDown:
				userInput <- game.ActionDown
			case termbox.KeyEnter:
				userInput <- game.ActionDown
			case termbox.KeyArrowLeft:
				userInput <- game.ActionLeft
			case termbox.KeyArrowRight:
				userInput <- game.ActionRight
			}
		}
	}()

	go func() {
		for state := range tetris.Updates() {
			w, h := termbox.Size()
			if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
				panic(err)
			}

			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					termbox.SetCell(x, y, ' ', termbox.ColorDefault,
						termbox.Attribute(rand.Int()%8)+1)
				}
			}
			if err := termbox.Flush(); err != nil {
				panic(err)
			}
			_ = state
		}
	}()

	tetris.Init()
	go tetris.MainLoop()

	select {
	case <-tetris.Wait():
	case <-signals:
	}

}
