package game

import (
	"tetris/src/shape"
)

const (
	FieldWith   = 10
	FieldHeight = 20
)

type Field [FieldHeight][FieldWith]Cell

func (fields *Field) Set(figure fallingFigure) {
	state := figure.GetCurrentState()
	for point := range fields.indexGenerator(state, figure.point) {

		fields[point.Y][point.X] = Cell{
			Cell:  shape.Cell{Filled: true},
			Color: figure.Color,
		}
	}
}

func (fields *Field) Remove(figure fallingFigure) {
	state := figure.GetCurrentState()
	for point := range fields.indexGenerator(state, figure.point) {

		fields[point.Y][point.X] = Cell{}
	}
}

func (fields *Field) indexGenerator(shapeState shape.State, shift point) <-chan point {
	ch := make(chan point)

	go func() {
		for y := range shapeState {
			for x := range shapeState[y] {
				if !shapeState[y][x].Filled {
					continue
				}

				point := point{
					X: x + shift.X,
					Y: y + shift.Y,
				}

				if !fields.inBounds(point) {
					continue
				}

				ch <- point
			}
		}
		close(ch)
	}()

	return ch
}

func (Field) inBounds(index point) bool {
	return index.Y >= 0 && index.Y < FieldHeight && index.X >= 0 && index.X < FieldWith
}
