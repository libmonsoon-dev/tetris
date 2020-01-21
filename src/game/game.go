package game

type Game struct {
	ticker        chan struct{}
	close         chan struct{}
	actions       chan Action
	updates       chan Snapshot
	pause         chan struct{}
	state         State
	fallingFigure fallingFigure
}

func New() *Game {
	game := &Game{
		close:   make(chan struct{}, closeChanCap),
		actions: make(chan Action, actionsChanCap),
		updates: make(chan Snapshot, updatesChanCap),
		pause:   make(chan struct{}),
		state: State{
			FullField: FullField{},
			Score:     0,
			Next:      RandomShape(),
		},
	}
	game.newFallingFigure()

	return game
}

func (game Game) Actions() chan<- Action {
	return game.actions
}

func (game Game) Updates() <-chan Snapshot {
	return game.updates
}

func (game *Game) Init() {
	game.initTicker()
}

func (game *Game) MainLoop() {

	for {
		select {
		case <-game.close:
			//TODO: render game over screen
			println("Game over! Score:", game.state.Score)
			return
		case action := <-game.actions:
			game.processAction(action)
		case <-game.ticker:
			game.processNextStep()
		}
		game.updateState()
	}
}

func (game Game) Wait() <-chan struct{} {
	return game.close
}

func (game Game) Exit() {
	close(game.close)
}
