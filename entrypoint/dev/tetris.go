// +build dev

package main

import (
	"os"
	"os/signal"
	"syscall"

	"tetris/pkg/game"

	"github.com/nsf/termbox-go"
)

type IO struct {
	width  int
	height int
}

func (io *IO) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	io.updateSize()
}

func (io IO) Close() {
	termbox.Close()
}

func (io *IO) Render(snapshot game.Snapshot) {
	io.clear()
	io.renderBox()
	io.renderLogo()
	io.renderField(snapshot.Field)
	io.renderScore(snapshot.Score)
	io.renderNextShape(snapshot.Next)
	io.flush()
}

func (io *IO) updateSize() {
	io.width, io.height = termbox.Size()
}

func (io *IO) clear() {
	if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
		panic(err)
	}
}

func (io *IO) flush() {
	if err := termbox.Flush(); err != nil {
		panic(err)
	}
}

func (io *IO) renderLogo() {

}

func (io *IO) renderBox() {
	const xShift = 5
	const yShift = 5
	const color = termbox.ColorWhite

	for i := -1; i < game.FieldHeight+1; i++ {
		io.setCell(xShift-1, yShift+i, color)
		io.setCell(xShift+game.FieldWith, yShift+i, color)
	}

	for i := -1; i < game.FieldWith+1; i++ {
		io.setCell(xShift+i, yShift-1, color)
		io.setCell(xShift+i, yShift+game.FieldHeight, color)
	}
}

func (io *IO) renderField(fields game.Field) {
	const xShift = 5
	const yShift = 5
	const emptyCellColor = termbox.ColorBlack

	for x := 0; x < game.FieldWith; x++ {
		for y := 0; y < game.FieldHeight; y++ {
			cell := fields[y][x]

			if cell.Filled {
				io.setCell(x+xShift, y+yShift, io.getColor(cell))
			} else {
				io.setCell(x+xShift, y+yShift, emptyCellColor)
			}
		}
	}
}

func (io *IO) setCell(x, y int, bg termbox.Attribute) {
	const ch = ' '

	termbox.SetCell(x*2, y, ch, termbox.ColorDefault, bg)
	termbox.SetCell(x*2+1, y, ch, termbox.ColorDefault, bg)
}

func (io *IO) getColor(cell game.Cell) termbox.Attribute {
	switch cell.Color {
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

func (io *IO) renderScore(i int) {

}

func (io *IO) renderNextShape(shape game.Shape) {

}

func main() {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	io := &IO{}

	io.Init()
	defer io.Close()

	tetris := game.DevRestore()
	defer tetris.DevDump()

	go func() {
		userInput := tetris.Actions()

		for {
			event := termbox.PollEvent()

			switch event.Key {
			case termbox.KeyEsc:
				tetris.Close()
			case termbox.KeySpace:
				userInput <- game.ActionPause
			case termbox.KeyArrowUp:
				userInput <- game.ActionUp
			case termbox.KeyArrowDown:
				userInput <- game.ActionDown
			case termbox.KeyEnter:
				userInput <- game.ActionDown
			case termbox.KeyArrowLeft:
				userInput <- game.ActionLeft
			case termbox.KeyArrowRight:
				userInput <- game.ActionRight
			}
		}
	}()

	go func() {
		for state := range tetris.Updates() {
			io.Render(state)
		}
	}()

	tetris.Init()
	go tetris.MainLoop()

	select {
	case <-tetris.Wait():
	case <-signals:
	}

}
