// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package game

import (
	"reflect"
	"testing"
)

func TestChoiceRandomColor(t *testing.T) {
	list := make([]Color, 1)
	result := choiceRandomColor(list)

	if !reflect.DeepEqual(result, list[0]) {
		t.Fatalf("!reflect.DeepEqual(%#v, %#v)", result, list[0])
	}
}

func TestChoiceRandomColorPanic(t *testing.T) {
	defer func() {
		err := recover()

		if !reflect.DeepEqual(err, choiceRandomPanicMsg) {
			t.Errorf("Invalid panic: expect \"%v\", got \"%v\"", choiceRandomPanicMsg, err)
		}
	}()

	choiceRandomColor(make([]Color, 0))
}