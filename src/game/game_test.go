package game

import (
	"testing"
)

func TestNew(t *testing.T) {
	game := New()

	if game.actions == nil {
		t.Error("game.actions == nil")
	}

	if game.updates == nil {
		t.Error("game.updates == nil")
	}

	if game.close == nil {
		t.Error("game.close == nil")
	}

	if val := len(game.actions); val != actionsChanLen {
		t.Errorf("len(game.actions) == %v", val)
	}

	if val := len(game.updates); val != updatesChanLen {
		t.Errorf("len(game.updates) == %v", val)
	}

	if val := len(game.close); val != closeChanLen {
		t.Errorf("len(game.close) == %v", val)
	}

	if val := cap(game.actions); val != actionsChanCap {
		t.Errorf("cap(game.actions) == %v", val)
	}

	if val := cap(game.updates); val != updatesChanCap {
		t.Errorf("cap(game.updates) == %v", val)
	}

	if val := cap(game.close); val != closeChanCap {
		t.Errorf("cap(game.close) == %v", val)
	}

}

func TestActions(t *testing.T) {
	game := New()
	want := game.actions

	if got := game.Actions(); got != want {
		t.Errorf("game.Actions() = %v, want %v", got, want)
	}
}

func TestUpdates(t *testing.T) {
	game := New()
	want := game.updates

	if got := game.Updates(); got != want {
		t.Errorf("game.Updates() = %v, want %v", got, want)
	}
}
