package game

import "time"

const (
	fallingFigureXShift = 3
	fallingFigureYShift = -5
)

func (game *Struct) processAction(action Action) {
	if !game.isValidAction(action) {
		return
	}
	game.doAction(action)
}

func (game Struct) isValidAction(action Action) bool {
	switch action {
	case ActionUp:
		return game.validateRotation()
	case ActionRight:
		return game.validateMove(1)
	case ActionLeft:
		return game.validateMove(-1)

	default:
		return true

	}
}

func (game *Struct) validateRotation() bool {
	game.state.Remove(game.fallingFigure)
	defer game.state.Set(game.fallingFigure)

	nextState := game.fallingFigure.Copy()
	nextState.Rotate()

	//TODO: refactor it
	return game.state.Field.CanBeSet(nextState)
}

func (game *Struct) validateMove(distance int) bool {
	nextState := game.fallingFigure.Copy()
	nextState.X += distance

	//TODO: refactor it
	return game.state.Field.CanBeSet(nextState)
}

func (game *Struct) doAction(action Action) {
	game.state.Remove(game.fallingFigure)

	switch action {
	case ActionPause:
		game.pauseSwitch()
	case ActionUp:
		game.fallingFigure.Rotate()
	case ActionRight:
		game.fallingFigure.Move(1)
	case ActionLeft:
		game.fallingFigure.Move(-1)
	}

	game.state.Set(game.fallingFigure)
}

func (game *Struct) processNextStep() {
	game.state.Remove(game.fallingFigure)

	game.fallingFigure.point.Y++
	if game.state.IsAtBottom(game.fallingFigure) {
		game.fallingFigure.point.Y--
		game.state.Set(game.fallingFigure)
		game.checkGameOver()
		game.clearLines()
		game.newFallingFigure()
	}
	game.state.Set(game.fallingFigure)
}

//TODO:
func (game *Struct) checkGameOver() {}

//TODO:
func (game *Struct) clearLines() {}

func (game *Struct) newFallingFigure() {
	game.fallingFigure = fallingFigure{
		Shape: game.state.Next,
		point: point{
			X: fallingFigureXShift,
			Y: fallingFigureYShift,
		},
	}
	game.state.Next = RandomShape()
}

func (game *Struct) pauseSwitch() {
	game.state.OnPause = !game.state.OnPause
	game.pause <- struct{}{}
}

func (game Struct) updateState() {
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

			case <-game.pause:
				select {
				case <-game.close:
					return
				case <-game.pause:
					//pass
				}
			}
		}
	}()
}
