package game

import (
	"tetris/src/shape"
)

const (
	FullFieldWith   = VisibleFieldWith
	FullFieldHeight = VisibleFieldHeight + 10
)

type FullField [FullFieldHeight][FullFieldWith]Cell

//TODO: review this file

func (field *FullField) Set(figure fallingFigure) {
	state := figure.GetCurrentState()
	for point := range field.indexGenerator(state, figure.point) {
		if !inBounds(point) {
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
		if !inBounds(point) {
			continue
		}

		field[point.Y][point.X] = Cell{}
	}
}

func (field *FullField) IsAtBottom(figure fallingFigure) bool {
	state := figure.GetCurrentState()
	for point := range field.indexGenerator(state, figure.point) {
		if isAbove(point) {
			continue
		}

		if isBelow(point) || field[point.Y][point.X].Filled {
			return true
		}
	}

	return false
}

func (field *FullField) CanBeSet(figure fallingFigure) bool {
	state := figure.GetCurrentState()
	for point := range field.indexGenerator(state, figure.point) {
		if !inBounds(point) {
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
	const heightDiff = FullFieldHeight - VisibleFieldHeight

	for i := range result {
		for j := range result[i] {
			result[i][j] = field[i+heightDiff][j]
		}
	}

	return
}

func isAbove(index point) bool {
	return index.Y < 0
}

func isBelow(index point) bool {
	return index.Y >= FullFieldHeight
}

func inBounds(index point) bool {
	return !isAbove(index) && !isBelow(index) && index.X >= 0 && index.X < FullFieldWith
}
