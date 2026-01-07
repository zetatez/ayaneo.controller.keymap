package main

import (
	"time"
	"log"

	evdev "github.com/gvalkov/golang-evdev"
)

func loop(cfg *Config, ui *UInput) error {
	dev, err := evdev.Open(cfg.Device)
	if err != nil {
		return err
	}

	log.Println("Opened device:", dev.Name)

	for {
		events, err := dev.Read()
		if err != nil {
			time.Sleep(5 * time.Millisecond)
			continue
		}

		for _, e := range events {
			// 🔴 核心调试日志
			log.Printf(
				"EV type=%d code=%d value=%d",
				e.Type, e.Code, e.Value,
			)

			switch e.Type {
			case evdev.EV_KEY:
				handleButton(ui, cfg, e)
			case evdev.EV_ABS:
				handleAxis(ui, cfg, e)
			}
		}
	}
}

