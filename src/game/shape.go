package game

import "tetris/src/shape"

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

func (sh *Shape) NextState() {
	sh.Index++
	sh.Index %= len(sh.Shape)
}
