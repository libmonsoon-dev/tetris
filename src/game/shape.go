package game

import (
	"tetris/src/shape"
)

type Shape struct {
	shape.Shape
	Index int
	Color
}

func RandomShape() Shape {
	sh, index := shape.ChoiceRandom()

	return Shape{
		Shape: sh,
		Index: index,
		Color: RandomColor(),
	}

}

func (sh Shape) GetCurrentState() shape.State {
	return sh.Shape[sh.Index]
}

func (sh Shape) Copy() Shape {
	return sh
}

func (sh *Shape) Rotate() {
	sh.Index = sh.getNextIndex()
}

func (sh Shape) getNextIndex() int {
	index := sh.Index
	return (index + 1) % len(sh.Shape)
}
