package shape

import "math/rand"

func ChoiceRandom() (Shape, int) {
	shape := choiceRandomShape(List[:])

	return shape, rand.Intn(len(shape))
}

var List = [...]Shape{S, Z, I, O, J, L, T}
