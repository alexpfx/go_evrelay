package evdev

import (
	"time"
)

const (
	KeyPress   KeyEventType = 0
	KeyRelease KeyEventType = 1
	KeyHold    KeyEventType = 2
)

type KeyEventType int32

type KeyEvent struct {
	Time time.Time
	Key  uint16
	Type KeyEventType
}

var KeyEventTypeMap = map[KeyEventType]string{
	KeyPress:   "KeyPress",
	KeyHold:    "KeyHold",
	KeyRelease: "KeyRelease",
}

func NewKeyEvent(rawEvent InputEvent) KeyEvent {
	return KeyEvent{
		Time: time.Now(),
		Key:  rawEvent.Code,
		Type: KeyEventType(rawEvent.Value),
	}
}

func (d Device) ReadKeys() (chan KeyEvent, chan bool) {
	ch := make(chan KeyEvent)
	qch := make(chan bool)
	go func() {
		for {
			rawEvents := d.read()

			for _, rev := range rawEvents {
				re := rev
				if re.Type != EvKey {
					continue
				}

				ch <- NewKeyEvent(re)
			}

			select {
			case stop := <-qch:

				d.Close()
				close(ch)

			default:

			}

		}

	}()

	return ch, qch
}
