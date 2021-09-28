package main

import (
	"log"

	"github.com/caius/z4pi/internal/app"
)

func main() {
	// FIXME: turn these into cli arguments
	brain, err := app.NewApp("/dev/ttyUSB0", 8080)
	if err != nil {
		log.Fatal(err)
	}

	err = brain.Run()
	if err != nil {
		log.Fatal(err)
	}
}
