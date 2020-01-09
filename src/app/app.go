package app

import (
	"tetris/src/game"
)

type app struct {
	ui   UI
	game *game.Struct
	//TODO: add errorHandler
}

func Init(ui UI, game *game.Struct) app {
	app := app{
		ui:   ui,
		game: game,
	}

	app.init()

	return app

}

func (app app) init() {
	app.ui.Init()
	app.game.Init()

	go func() {
		actions := app.game.Actions()

		for action := range app.ui.Actions() {
			if action == game.ActionExit {
				app.game.Exit()
				return
			} else {
				actions <- action
			}
		}
	}()

	go func() {
		for state := range app.game.Updates() {
			app.ui.Render(state)
		}
	}()
}

func (app app) Close() {
	app.ui.Close()
}

func (app app) MainLoop() {
	app.game.MainLoop()
}

func (app app) Wait() <-chan struct{} {
	return app.game.Wait()
}
