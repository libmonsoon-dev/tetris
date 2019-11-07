package slice

import (
	"reflect"
	"testing"
)

func TestChoiceRandomΔ(t *testing.T) {
	list := make([]Δ, 1)
	result := choiceRandomΔ(list)

	if !reflect.DeepEqual(result, list[0]) {
		t.Fatalf("!reflect.DeepEqual(%#v, %#v)", result, list[0])
	}
}

func TestChoiceRandomΔPanic(t *testing.T) {
	defer func() {
		err := recover()

		if !reflect.DeepEqual(err, choiceRandomPanicMsg) {
			t.Errorf("Invalid panic: expect \"%v\", got \"%v\"", choiceRandomPanicMsg, err)
		}
	}()

	choiceRandomΔ(make([]Δ, 0))
}
