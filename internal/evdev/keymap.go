package evdev

type keyMap struct {

}

func (k keyMap) GetName(code uint16) string {
	kmap := *km
	return kmap[code]
}

func (k keyMap) GetCode(name string) uint16 {
	return 0
}

func NewKeyMap() KeyMap{
	return &keyMap{}
}

type KeyMap interface {
	GetName(code uint16) string
	GetCode(name string) uint16
}