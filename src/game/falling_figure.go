package game

type fallingFigure struct {
	point
	Shape
}

func (f fallingFigure) Copy() fallingFigure {
	return f
}

func (f *fallingFigure) Move(distance int) {
	f.X += distance
}
