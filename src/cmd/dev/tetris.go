// +build dev

package main

import (
	"os"
	"os/signal"
	"syscall"

	"tetris/src/app"
	"tetris/src/console"
	"tetris/src/game"
)

func main() {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	logicComponent := game.DevRestore()
	defer logicComponent.DevDump()

	tetris := app.Init(
		&console.UI{},
		logicComponent,
	)
	defer tetris.Close()

	go tetris.MainLoop()

	select {
	case <-tetris.Wait():
	case <-signals:
	}

}
