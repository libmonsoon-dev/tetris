package game

type Snapshot struct {
	Field
	Score   int
	Next    Shape
	OnPause bool
}
