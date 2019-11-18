package game

import (
	"time"
)

type Struct struct {
	ticker  chan struct{}
	close   chan struct{}
	actions chan Action
	updates chan Snapshot
	state   Snapshot
}

func New() *Struct {
	return &Struct{
		close:   make(chan struct{}, closeChanCap),
		actions: make(chan Action, actionsChanCap),
		updates: make(chan Snapshot, updatesChanCap),
	}
}

func (game Struct) Actions() chan<- Action {
	return game.actions
}

func (game Struct) Updates() <-chan Snapshot {
	return game.updates
}

func (game *Struct) Init() {
	game.initTicker()
}

func (game *Struct) MainLoop() {

	for {
		select {
		case <-game.close:
			return
		case action := <-game.actions:
			game.processAction(action)
		case <-game.ticker:
			game.processNextStep()
		}
	}
}

func (game Struct) processAction(action Action) {
	if !game.isValidAction(action) {
		return
	}
	game.doAction(action)
	game.updates <- game.state
}

func (game Struct) isValidAction(action Action) bool {
	return true
}

func (game *Struct) doAction(action Action) {
}

func (game *Struct) processNextStep() {
	game.updates <- game.state
}

func (game *Struct) initTicker() {
	game.ticker = make(chan struct{})

	go func() {
		const waitInterval = initWaitInterval

		for {
			select {
			case <-game.close:
				return
			case game.ticker <- struct{}{}:
				time.Sleep(waitInterval)
			}
		}
	}()
}

func (game Struct) Wait() <-chan struct{} {
	return game.close
}

func (game Struct) Exit() {
	close(game.close)
}