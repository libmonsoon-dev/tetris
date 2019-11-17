package slice

import (
	"math/rand"

	"github.com/cheekybits/genny/generic"
)

type Δ generic.Type

const choiceRandomPanicMsg = "Impossible to get item from empty slice"

func choiceRandomΔ(list []Δ) Δ {
	listLength := len(list)

	if listLength == 0 {
		panic(choiceRandomPanicMsg)
	}

	index := rand.Intn(listLength)
	return list[index]

}
