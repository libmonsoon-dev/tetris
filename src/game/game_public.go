package game

type Struct struct {
	ticker        chan struct{}
	close         chan struct{}
	actions       chan Action
	updates       chan Snapshot
	pause         chan struct{}
	state         Snapshot
	fallingFigure fallingFigure
}

func New() *Struct {
	game := &Struct{
		close:   make(chan struct{}, closeChanCap),
		actions: make(chan Action, actionsChanCap),
		updates: make(chan Snapshot, updatesChanCap),
		pause:   make(chan struct{}),
		state: Snapshot{
			Field: Field{},
			Score: 0,
			Next:  RandomShape(),
		},
	}
	game.newFallingFigure()

	return game
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
		game.updateState()
	}
}

func (game Struct) Wait() <-chan struct{} {
	return game.close
}

func (game Struct) Exit() {
	close(game.close)
}
