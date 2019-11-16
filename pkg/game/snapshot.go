package game

const (
	FieldWith   = 10
	FieldHeight = 20
)

type Snapshot struct {
	Field [FieldHeight][FieldWith]bool
	Score int
}
