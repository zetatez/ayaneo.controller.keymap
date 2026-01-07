package main

import (
	"log"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	ui, err := NewUInput()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("AYANEO key mapper started")
	if err := loop(cfg, ui); err != nil {
		log.Fatal(err)
	}
}
