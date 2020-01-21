package game

type fallingFigure struct {
	point
	Shape
}

func (f fallingFigure) Copy() fallingFigure {
	return f
}
