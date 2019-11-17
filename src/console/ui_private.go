package console

import (
	"github.com/nsf/termbox-go"

	"tetris/src/game"
)

func (ui *UI) renderLogo() {

}

func (ui *UI) renderBox() {
	const xShift = 5
	const yShift = 5
	const color = termbox.ColorWhite

	for i := -1; i < game.FieldHeight+1; i++ {
		ui.setCell(xShift-1, yShift+i, color)
		ui.setCell(xShift+game.FieldWith, yShift+i, color)
	}

	for i := -1; i < game.FieldWith+1; i++ {
		ui.setCell(xShift+i, yShift-1, color)
		ui.setCell(xShift+i, yShift+game.FieldHeight, color)
	}
}

func (ui *UI) renderField(fields game.Field) {
	const xShift = 5
	const yShift = 5
	const emptyCellColor = termbox.ColorBlack

	for x := 0; x < game.FieldWith; x++ {
		for y := 0; y < game.FieldHeight; y++ {
			cell := fields[y][x]

			if cell.Filled {
				ui.setCell(x+xShift, y+yShift, ui.getColor(cell.Color))
			} else {
				ui.setCell(x+xShift, y+yShift, emptyCellColor)
			}
		}
	}
}

func (ui *UI) renderScore(i int) {

}

func (ui *UI) renderNextShape(shape game.Shape) {
	const xShift = 20
	const yShift = 5

	state := shape.GetCurrentState()

	for y := range state {
		for x := range state[y] {
			if state[y][x].Filled {
				ui.setCell(x+xShift, y+yShift, ui.getColor(shape.Color))
			}
		}
	}

}

func (ui *UI) setCell(x, y int, bg termbox.Attribute) {
	const ch = ' '

	termbox.SetCell(x*2, y, ch, termbox.ColorDefault, bg)
	termbox.SetCell(x*2+1, y, ch, termbox.ColorDefault, bg)
}

func (UI) getColor(color game.Color) termbox.Attribute {
	switch color {
	case game.ColorBlack:
		return termbox.ColorBlack
	case game.ColorRed:
		return termbox.ColorRed
	case game.ColorGreen:
		return termbox.ColorGreen
	case game.ColorYellow:
		return termbox.ColorYellow
	case game.ColorBlue:
		return termbox.ColorBlue
	case game.ColorMagenta:
		return termbox.ColorMagenta
	case game.ColorCyan:
		return termbox.ColorCyan
	case game.ColorWhite:
		return termbox.ColorWhite
	default:
		return termbox.ColorDefault
	}
}

func (ui *UI) updateSize() {
	ui.width, ui.height = termbox.Size()
}

func (ui *UI) clear() error {
	return termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func (ui *UI) flush() error {
	return termbox.Flush()
}
