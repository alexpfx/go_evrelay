package evdev

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getEventPaths() []string {
	res := make([]string, 0)
	baseDir := "/dev/input"

	root, err := os.Open(baseDir)
	if err != nil {
		log.Fatal(err)
	}

	defer root.Close()
	names, err := root.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		if !strings.HasPrefix(name, "event") {
			continue
		}
		path := filepath.Join(baseDir, name)
		if !isInputDevice(path) {
			continue
		}

		res = append(res, path)
	}

	return res
}

func openEventFile(path string) (*os.File, error) {
	var err error
	if !isInputDevice(path) {
		return nil, errors.New("not a input device: " + path)
	}

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return f, nil
}

func isBitSet(bitPos, number byte) bool {
	return 1<<bitPos&number != 0
}

func isKeyboard(capMap map[uint16][]uint16) bool {
	if hasAnyEventType(capMap, EvRel, EvAbs, EvFf, EvPwr) {
		return false
	}

	return hasAnyEventType(capMap, EvKey) &&
		hasAnyEventCode(capMap, EvKey, KeyA, KeyKp1)
}

func extractVersion(v int32) []int {
	r := make([]int, 3)
	r[0] = int(v >> 16)
	r[1] = int(v >> 8 & 0xff)
	r[2] = int(v >> 0 & 0xff)
	return r
}

func isInputDevice(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	m := f.Mode()
	return m&os.ModeCharDevice != 0
}

func hasAnyEventCode(codeMap map[uint16][]uint16, codeType uint16, codeTest ...uint16) bool {
	var codes []uint16
	var ok bool
	if codes, ok = codeMap[codeType]; !ok {
		return false
	}

	for _, c := range codes {
		for _, t := range codeTest {
			if c == t {
				return true
			}
		}
	}
	return false
}
func hasAnyEventType(capMap map[uint16][]uint16, capsTest ...uint16) bool {
	for _, cpt := range capsTest {
		_, ok := capMap[cpt]
		if ok {
			return true
		}
	}
	return false
}
