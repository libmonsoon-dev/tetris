package main

import (
	"fmt"
	"os"
	"time"
)

func clearScreen() {
	_, err := os.Stdout.Write([]byte{0x1B, 0x5B, 0x33, 0x3B, 0x4A, 0x1B, 0x5B, 0x48, 0x1B, 0x5B, 0x32, 0x4A})

	if err != nil {
		panic(err)
	}
}

func showTime(t time.Time) {
	fmt.Println(t)
}

func main() {
	var frame byte
	for now := range time.Tick(time.Second / 60) {
		frame++
		if frame%15 != 0 {
			continue
		}

		clearScreen()
		showTime(now)
	}
}
