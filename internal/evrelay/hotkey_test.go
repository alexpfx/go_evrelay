package evrelay

import (
	"github.com/alexpfx/go_evrelay/internal/evdev"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHotkey_KeyMatch(t *testing.T) {
	seq1 := []uint16{evdev.KeyA, evdev.KeyB, evdev.KeyL, evdev.Key0, evdev.Key1}

	table := []struct {
		name   string
		hotkey Hotkey
	}{
		{name: "empty", hotkey: Hotkey{
			Mods:   []uint16{},
			KeySeq: []uint16{},
		}},
		{name: "1 key and no match", hotkey: Hotkey{
			Mods:   []uint16{},
			KeySeq: []uint16{evdev.Key8},
		}},
		{name: "1 key and match first", hotkey: Hotkey{
			Mods:   []uint16{},
			KeySeq: []uint16{evdev.KeyA},
		}},
		{name: "1 key and match nth", hotkey: Hotkey{
			Mods:   []uint16{},
			KeySeq: []uint16{evdev.Key1},
		}},
		{name: "3 key and match ", hotkey: Hotkey{
			Mods:   []uint16{},
			KeySeq: []uint16{evdev.KeyL, evdev.Key0,evdev.Key1},
		}},
	}

	for _, ts := range table {
		action := func(h *Hotkey) {
			assert.EqualValues(t, h.KeySeq, ts.hotkey.KeySeq)
		}

		test := ts
		t.Run(test.name, func(t *testing.T) {
			for _, k := range seq1 {
				test.hotkey.KeyMatch(k, &action)
			}
		})
	}

}
