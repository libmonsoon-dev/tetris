package app

import (
	"tetris/src/game"
)

type App struct {
	ui   UI
	game *game.Struct
	//TODO: add errorHandler
}

func New(ui UI, game *game.Struct) App {
	return App{ui: ui, game: game}
}

func (app App) Init() {
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

func (app App) Close() {
	app.ui.Close()
}

func (app App) MainLoop() {
	app.game.MainLoop()
}

func (app App) Wait() <-chan struct{} {
	return app.game.Wait()
}
