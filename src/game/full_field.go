package game

import (
	"tetris/src/shape"
)

const (
	FullFieldWith   = VisibleFieldWith
	FullFieldHeight = VisibleFieldHeight + 10
	heightDiff      = FullFieldHeight - VisibleFieldHeight
)

type FullField [FullFieldHeight][FullFieldWith]Cell

//TODO: review this file

func (field *FullField) Set(figure fallingFigure) {
	state := figure.GetCurrentState()
	for point := range field.indexGenerator(state, figure.point) {
		if !inFullFieldBounds(point) {
			continue
		}

		field[point.Y][point.X] = Cell{
			Cell:  shape.Cell{Filled: true},
			Color: figure.Color,
		}
	}
}

func (field *FullField) Remove(figure fallingFigure) {
	state := figure.GetCurrentState()
	for point := range field.indexGenerator(state, figure.point) {
		if !inFullFieldBounds(point) {
			continue
		}

		field[point.Y][point.X] = Cell{}
	}
}

func (field *FullField) IsAtBottom(figure fallingFigure) bool {
	state := figure.GetCurrentState()
	for point := range field.indexGenerator(state, figure.point) {
		if isAboveField(point) {
			continue
		}

		point.Down()
		if isBelowFullField(point) || field[point.Y][point.X].Filled {
			return true
		}
	}

	return false
}

func (field *FullField) CanBeSet(figure fallingFigure) bool {
	state := figure.GetCurrentState()
	for point := range field.indexGenerator(state, figure.point) {
		if !inFullFieldBounds(point) {
			return false
		}
	}

	return true
}

func (FullField) indexGenerator(shapeState shape.State, shift point) <-chan point {
	ch := make(chan point)

	go func() {
		for y := range shapeState {
			for x := range shapeState[y] {
				if !shapeState[y][x].Filled {
					continue
				}

				ch <- point{
					X: x + shift.X,
					Y: y + shift.Y,
				}
			}
		}
		close(ch)
	}()

	return ch
}

func (field *FullField) getVisiblePath() (result VisibleField) {
	for i := range result {
		for j := range result[i] {
			result[i][j] = field[i+heightDiff][j]
		}
	}

	return
}

func (field *FullField) haveActiveInvisibleCells() bool {
	for i := 0; i < heightDiff; i++ {
		for j := range field[i] {
			if field[i][j].Filled {
				return true
			}
		}
	}

	return false
}
