// +build darwin linux windows

package main

import (
    "fmt"
    "golang.org/x/mobile/app"
)

func main() {
    app.Main(func(a app.App) {
        for event := range a.Events() {
            fmt.Println(event)
        }
    })
}
