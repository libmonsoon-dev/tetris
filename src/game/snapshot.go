package game

type Snapshot struct {
	VisibleField
	Score   int
	Next    Shape
	OnPause bool
}
