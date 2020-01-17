package game

type State struct {
	FullField
	Score   int
	Next    Shape
	OnPause bool
}

func (s State) GetSnapshot() Snapshot {
	return Snapshot{
		s.FullField.getVisiblePath(),
		s.Score,
		s.Next,
		s.OnPause,
	}

}
