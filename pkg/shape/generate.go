package shape

//go:generate genny -in=../../pkg/utils/genny_template/slice/choice_item.go -out=gen_choice_shape.go -pkg=shape gen "Δ=Shape"
//go:generate genny -in=../../pkg/utils/genny_template/slice/choice_item_test.go -out=gen_choice_shape_test.go -pkg=shape gen "Δ=Shape"
