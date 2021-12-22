package main

import (
	"log"

	"nyiyui.ca/koku/cmd/koku/internal"
)

func main() {
	err := internal.Main()
	if err != nil {
		log.Fatal(err)
	}
}
