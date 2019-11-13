package game

type Struct struct {
	Close   chan struct{}
	actions chan Action
	updates chan Snapshot
	score   int
	state   Snapshot
}

func New() *Struct {
	game := &Struct{
		Close:   make(chan struct{}, closeChanCap),
		actions: make(chan Action, actionsChanCap),
		updates: make(chan Snapshot, updatesChanCap),
		score:   initScore,
	}

	go game.processUserInput()

	return game
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

func (game *Struct) processUserInput() {
	for {
		select {
		case <-game.Close:
			return
		case action := <-game.actions:
			game.processAction(action)
		}
	}
}

func (game Struct) processAction(action Action) {
	if !game.validAction(action) {
		return
	}
	game.DoAction(action)
}

func (game Struct) validAction(action Action) bool {
	//TODO
	panic("Implement me")
	return true
}

func (game *Struct) DoAction(action Action) {
	//TODO
	panic("Implement me")
}
