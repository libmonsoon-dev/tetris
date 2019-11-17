package shape

type list []Shape

func (l list) ChoiceRandom() Shape {
	return choiceRandomShape(l)
}

var List = list{S, Z, I, O, J, L, T}
