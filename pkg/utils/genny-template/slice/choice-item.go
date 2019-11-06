package slice

import (
	"math/rand"

	"github.com/cheekybits/genny/generic"
)

type T generic.Type

func choiceRandomT(list []T) T {
	listLength := len(list)

	if listLength == 0 {
		panic("Impossible to get item from empty slice")
	}

	index := rand.Intn(listLength)
	return list[index]

}
