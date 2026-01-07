package main

import (
	evdev "github.com/gvalkov/golang-evdev"
)

func handleButton(ui *UInput, cfg *Config, e evdev.InputEvent) {
	name := evdev.KEY[int(e.Code)]
	if name == "" {
		return
	}

	action, ok := cfg.Buttons[name]
	if !ok {
		return
	}

	if e.Value == 1 {
		switch v := action.(type) {
		case string:
			ui.SendKey(keyCode(v), 1)
		case map[string]interface{}:
			if combo, ok := v["combo"].([]interface{}); ok {
				for _, k := range combo {
					ui.SendKey(keyCode(k.(string)), 1)
				}
			}
		}
	} else if e.Value == 0 {
		switch v := action.(type) {
		case string:
			ui.SendKey(keyCode(v), 0)
		case map[string]interface{}:
			if combo, ok := v["combo"].([]interface{}); ok {
				for i := len(combo) - 1; i >= 0; i-- {
					ui.SendKey(keyCode(combo[i].(string)), 0)
				}
			}
		}
	}
}

func handleAxis(ui *UInput, cfg *Config, e evdev.InputEvent) {
	axis := evdev.ABS[int(e.Code)]
	m, ok := cfg.Axes[axis]
	if !ok {
		return
	}

	if e.Value < -cfg.Deadzone {
		ui.SendKey(keyCode(m.Negative), 1)
		ui.SendKey(keyCode(m.Positive), 0)
	} else if e.Value > cfg.Deadzone {
		ui.SendKey(keyCode(m.Positive), 1)
		ui.SendKey(keyCode(m.Negative), 0)
	} else {
		ui.SendKey(keyCode(m.Negative), 0)
		ui.SendKey(keyCode(m.Positive), 0)
	}
}

