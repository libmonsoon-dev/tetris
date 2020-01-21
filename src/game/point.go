package game

type point struct {
	X, Y int
}

func (f *point) Move(distance int) {
	f.X += distance
}

func (f *point) Up() {
	f.Y--
}

func (f *point) Down() {
	f.Y++
}

func isAboveField(index point) bool {
	return index.Y < 0
}

func isBelowFullField(index point) bool {
	return index.Y >= FullFieldHeight
}

func inFullFieldBounds(index point) bool {
	return !isAboveField(index) && !isBelowFullField(index) && index.X >= 0 && index.X < FullFieldWith
}
