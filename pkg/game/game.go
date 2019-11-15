package game

import (
	"time"
)

type Struct struct {
	ticker  chan struct{}
	Close   chan struct{}
	actions chan Action
	updates chan Snapshot
	score   int
	state   Snapshot
}

func New() *Struct {
	return &Struct{
		Close:   make(chan struct{}, closeChanCap),
		actions: make(chan Action, actionsChanCap),
		updates: make(chan Snapshot, updatesChanCap),
		score:   initScore,
	}
}

func (game Struct) Actions() chan<- Action {
	return game.actions
}

func (game Struct) Updates() <-chan Snapshot {
	return game.updates
}

func (game Struct) GetScore() int {
	return game.score
}

func (game *Struct) Init() {
	game.initTicker()
}

func (game *Struct) MainLoop() {

	for {
		select {
		case <-game.Close:
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
	panic("Implement me") // TODO
}

func (game *Struct) doAction(action Action) {
	panic("Implement me") // TODO
}

func (game *Struct) processNextStep() {
	game.updates <- game.state
	panic("Implement me") // TODO
}

func (game *Struct) initTicker() {
	game.ticker = make(chan struct{})

	go func() {
		const waitInterval = initWaitInterval

		for {
			select {
			case <-game.Close:
				return
			case game.ticker <- struct{}{}:
				time.Sleep(waitInterval)
			}
		}
	}()
}
