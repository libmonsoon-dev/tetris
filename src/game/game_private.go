package game

import "time"

const (
	fallingFigureXShift = 3
	fallingFigureYShift = -5
)

func (game Struct) processAction(action Action) {
	if !game.isValidAction(action) {
		return
	}
	game.doAction(action)
	game.updates <- game.state
}

func (Struct) isValidAction(action Action) bool {
	return true
}

func (game *Struct) doAction(action Action) {
}

func (game *Struct) processNextStep() {
	if game.fallingFigure == nil {
		game.newFallingFigure()
	}
	game.state.Remove(*game.fallingFigure)
	game.fallingFigure.point.Y++
	if game.state.IsAtBottom(*game.fallingFigure) {
		game.fallingFigure.point.Y--
		game.state.Set(*game.fallingFigure)
		game.fallingFigure = nil
	} else {
		game.state.Set(*game.fallingFigure)
	}
	game.updates <- game.state
}

func (game *Struct) newFallingFigure() {
	game.fallingFigure = &fallingFigure{
		Shape: game.state.Next,
		point: point{
			X: fallingFigureXShift,
			Y: fallingFigureYShift,
		},
	}
	game.state.Next = RandomShape()
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
