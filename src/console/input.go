package console

import (
	"github.com/nsf/termbox-go"

	"tetris/src/game"
)

func (ui UI) poolEventLoop() {
	for {
		switch termbox.PollEvent().Key {
		case termbox.KeyEsc:
			ui.actions <- game.ActionExit
		case termbox.KeySpace:
			ui.actions <- game.ActionPause
		case termbox.KeyArrowUp:
			ui.actions <- game.ActionUp
		case termbox.KeyArrowDown:
			ui.actions <- game.ActionDown
		case termbox.KeyEnter:
			ui.actions <- game.ActionDown
		case termbox.KeyArrowLeft:
			ui.actions <- game.ActionLeft
		case termbox.KeyArrowRight:
			ui.actions <- game.ActionRight
		}
	}
}
