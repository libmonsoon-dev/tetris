// +build dev

package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	fileName = "state.json"
	filePerm = 0666
)

type jsonDto struct {
	State Snapshot
}

func mustNil(e error) {
	if e != nil {
		panic(e)
	}

}

func (game Game) DevDump() {
	blob, err := json.Marshal(game)
	mustNil(err)

	err = ioutil.WriteFile(fileName, blob, filePerm)
	mustNil(err)
}

func DevRestore() (game *Game) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	game = New()

	blob, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(fmt.Errorf("game.DevRestore: ioutil.ReadFile(%v): %w", fileName, err))
		return
	}

	tmp := &Game{}
	err = json.Unmarshal(blob, tmp)

	if err != nil {
		fmt.Println(fmt.Errorf("game.DevRestore: %w", err))
		return
	}

	game.state = tmp.state
	return
}

func (game *Game) UnmarshalJSON(blob []byte) error {
	dto := &jsonDto{}
	err := json.Unmarshal(blob, dto)
	if err != nil {
		return err
	}

	game.state = dto.State

	return nil
}

func (game Game) MarshalJSON() ([]byte, error) {
	dto := jsonDto{State: game.state}

	return json.MarshalIndent(dto, "", strings.Repeat(" ", 4))
}
