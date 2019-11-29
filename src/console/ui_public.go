package console

import (
	"github.com/nsf/termbox-go"

	"tetris/src/game"
)

type UI struct {
	width  int
	height int
}

func (ui *UI) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	ui.updateSize()
}

func (ui UI) Close() {
	termbox.Close()
}

func (ui *UI) Render(snapshot game.Snapshot) error {
	if err := ui.clear(); err != nil {
		return err
	}

	ui.renderBox()
	ui.renderLogo()
	ui.renderField(snapshot.Field)
	ui.renderScore(snapshot.Score)
	ui.renderNextShape(snapshot.Next)
	if snapshot.OnPause {
		ui.renderPauseBar()
	}

	return ui.flush()
}

func (ui *UI) PoolAction() game.Action {
	for {
		switch termbox.PollEvent().Key {
		case termbox.KeyEsc:
			return game.ActionExit
		case termbox.KeySpace:
			return game.ActionPause
		case termbox.KeyArrowUp:
			return game.ActionUp
		case termbox.KeyArrowDown:
			return game.ActionDown
		case termbox.KeyEnter:
			return game.ActionDown
		case termbox.KeyArrowLeft:
			return game.ActionLeft
		case termbox.KeyArrowRight:
			return game.ActionRight
		}
	}
}
