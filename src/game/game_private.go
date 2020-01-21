package game

import (
	"time"
)

const (
	fallingFigureXShift = 3
	fallingFigureYShift = -5
)

func (game *Game) processAction(action Action) {
	if !game.isValidAction(action) {
		return
	}
	game.doAction(action)
}

func (game Game) isValidAction(action Action) bool {
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

func (game *Game) validateRotation() bool {
	game.state.Remove(game.fallingFigure)
	defer game.state.Set(game.fallingFigure)

	nextState := game.fallingFigure.Copy()
	nextState.Rotate()

	return game.state.FullField.CanBeSet(nextState)
}

func (game *Game) validateMove(distance int) bool {
	nextState := game.fallingFigure.Copy()
	nextState.X += distance

	return game.state.FullField.CanBeSet(nextState)
}

func (game *Game) doAction(action Action) {
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

func (game *Game) processNextStep() {
	game.hideFallingFigure()
	if !game.isFallingFigureAtBottom() {
		game.fallingFigure.Down()
		game.writeFallingFigure()
		return
	}

	game.writeFallingFigure() // we fix the figure, it no longer falls and now is a static part of the field
	game.checkGameOver()
	game.clearLines()
	game.newFallingFigure()

	game.writeFallingFigure()
}

func (game *Game) isFallingFigureAtBottom() bool {
	return game.state.IsAtBottom(game.fallingFigure)
}

func (game *Game) hideFallingFigure() {
	game.state.Remove(game.fallingFigure)
}

func (game *Game) writeFallingFigure() {
	game.state.Set(game.fallingFigure)
}

func (game *Game) checkGameOver() {
	game.state.GameOver = game.state.FullField.haveActiveInvisibleCells()

	if game.state.GameOver {
		game.Exit()
	}
}

//TODO:
func (game *Game) clearLines() {}

func (game *Game) newFallingFigure() {
	game.fallingFigure = fallingFigure{
		Shape: game.state.Next,
		point: point{
			X: fallingFigureXShift,
			Y: fallingFigureYShift,
		},
	}
	game.state.Next = RandomShape()
}

func (game *Game) pauseSwitch() {
	game.state.OnPause = !game.state.OnPause
	game.pause <- struct{}{}
}

func (game Game) updateState() {
	game.updates <- game.state.GetSnapshot()
}

func (game *Game) initTicker() {
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
