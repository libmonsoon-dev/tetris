package game

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestNew(t *testing.T) {
	game := New()

	if game.actions == nil {
		t.Error("game.actions == nil")
	}

	if game.updates == nil {
		t.Error("game.updates == nil")
	}

	if game.Close == nil {
		t.Error("game.Close == nil")
	}

	if val := len(game.actions); val != actionsChanLen {
		t.Errorf("len(game.actions) == %v", val)
	}

	if val := len(game.updates); val != updatesChanLen {
		t.Errorf("len(game.updates) == %v", val)
	}

	if val := len(game.Close); val != closeChanLen {
		t.Errorf("len(game.Close) == %v", val)
	}

	if val := cap(game.actions); val != actionsChanCap {
		t.Errorf("cap(game.actions) == %v", val)
	}

	if val := cap(game.updates); val != updatesChanCap {
		t.Errorf("cap(game.updates) == %v", val)
	}

	if val := cap(game.Close); val != closeChanCap {
		t.Errorf("cap(game.Close) == %v", val)
	}

	if game.score != initScore {
		t.Errorf("game.score = %v", game.score)
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

func TestGetScore(t *testing.T) {
	game := New()
	game.score = rand.Intn(math.MaxInt32)

	want := game.score
	if got := game.GetScore(); got != want {
		t.Errorf("GetScore() = %v, want %v", got, want)
	}
}
