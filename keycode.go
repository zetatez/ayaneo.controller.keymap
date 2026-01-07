package main

import (
	"fmt"

	evdev "github.com/gvalkov/golang-evdev"
)

// keyCode 把 "KEY_ENTER" → linux keycode(uint16)
func keyCode(name string) uint16 {
	for code, keyName := range evdev.KEY {
		if keyName == name {
			return uint16(code)
		}
	}
	panic(fmt.Sprintf("unknown key: %s", name))
}

