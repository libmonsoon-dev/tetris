package game

type Action int

const (
	_ Action = iota
	ActionExit
	ActionPause
	ActionUp
	ActionDown
	ActionLeft
	ActionRight
)
