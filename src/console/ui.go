package console

import (
	"github.com/nsf/termbox-go"

	"tetris/src/game"
)

const actionChanSize = 1

type UI struct {
	width   int
	height  int
	actions chan game.Action
	close   chan struct{}
}

func (ui *UI) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	ui.updateSize()
	ui.actions = make(chan game.Action, actionChanSize)
	ui.close = make(chan struct{})
	go ui.poolEventLoop()
}

func (ui UI) Close() {
	close(ui.close)
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

func (ui UI) Actions() <-chan game.Action {
	return ui.actions
}
