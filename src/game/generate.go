package game

//go:generate genny -in=../utils/genny_template/slice/choice_item.go -out=gen_choice_color.go -pkg=game gen "Δ=Color"
//go:generate genny -in=../utils/genny_template/slice/choice_item_test.go -out=gen_choice_color_test.go -pkg=game gen "Δ=Color"
