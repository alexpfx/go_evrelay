package evrelay

type Hotkey struct {
	Mods   []uint16
	KeySeq []uint16
	pos    int
}

func (h *Hotkey) KeyMatch(key uint16, action *func(hotkey *Hotkey)) {
	seqSize := len(h.KeySeq)
	if seqSize == 0 {
		return
	}

	if h.KeySeq[h.pos] != key {
		h.reset()
		return
	}
	h.pos++
	if h.pos >= seqSize {
		h.reset()
		f := *action
		f(h)
	}
}
func (h *Hotkey) reset() {
	h.pos = 0
}
