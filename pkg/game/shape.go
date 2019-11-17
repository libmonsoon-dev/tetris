package game

import "tetris/pkg/shape"

type Shape struct {
	shape.Shape
	Index int
	Color
}

func (sh Shape) GetCurrentState() shape.State {
	return sh.Shape[sh.Index]
}

func (sh *Shape) NextState() {
	sh.Index++
	sh.Index %= len(sh.Shape)
}
