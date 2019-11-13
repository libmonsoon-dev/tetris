package main

import (
	"fmt"
	"tetris/pkg/game"

	"github.com/nsf/termbox-go"
)

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	tetris := game.New()
	defer close(tetris.Close)

	go func() {
		userInput := tetris.Actions()

		for {
			event := termbox.PollEvent()

			switch event.Key {
			case termbox.KeyEsc:
				close(tetris.Close)
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
			err = termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			fmt.Println(state)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	<-tetris.Close
	fmt.Println("Score: ", tetris.GetScore())

}
