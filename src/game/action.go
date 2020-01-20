package game

type Action int

//go:generate gostringer -type=Action
//go:generate stringer -type=Action

const (
	_ Action = iota
	ActionExit
	ActionPause
	ActionUp
	ActionDown
	ActionLeft
	ActionRight
)
