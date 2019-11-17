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

func (game Struct) DevDump() {
	blob, err := json.Marshal(game)
	mustNil(err)

	err = ioutil.WriteFile(fileName, blob, filePerm)
	mustNil(err)
}

func DevRestore() (game *Struct) {
	game = New()

	blob, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(fmt.Errorf("game.DevRestore: ioutil.ReadFile(%v): %v", fileName, err))
		return
	}

	tmp := &Struct{}
	err = json.Unmarshal(blob, tmp)

	if err != nil {
		fmt.Println(fmt.Errorf("game.DevRestore: %v", err))
		return
	}

	game.state = tmp.state
	return
}

func (game *Struct) UnmarshalJSON(blob []byte) error {
	dto := &jsonDto{}
	err := json.Unmarshal(blob, dto)
	if err != nil {
		return err
	}

	game.state = dto.State

	return nil
}

func (game Struct) MarshalJSON() ([]byte, error) {
	dto := jsonDto{State: game.state}

	return json.MarshalIndent(dto, "", strings.Repeat(" ", 4))
}
