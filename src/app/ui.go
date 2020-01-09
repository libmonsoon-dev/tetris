package app

import "tetris/src/game"

type UI interface {
	Init()
	Close()
	Render(snapshot game.Snapshot) error
	Actions() <-chan game.Action
}
