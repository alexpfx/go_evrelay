package evdev

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"unsafe"
)

type Device struct {
	Name         string
	Uniq         string
	Phys         string
	Prop         string
	Version      string
	InputId      InputId
	Capabilities map[uint16][]uint16
	file         *os.File
}

func (d *Device) Close() {
	_ = d.file.Close()
}

func (d *Device) read() []InputEvent {
	readCount := 16

	buffer := make([]byte, int(unsafe.Sizeof(InputEvent{}))*readCount)

	_, err := d.file.Read(buffer)
	if err != nil {
		log.Fatalf(err.Error())
	}

	events := make([]InputEvent, readCount)
	b := bytes.NewBuffer(buffer)
	err = binary.Read(b, binary.LittleEndian, &events)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for i, event := range events {
		if event.Time.Sec == 0 {
			return events[:i]
		}
	}

	return events
}

func getDeviceInfo(f *os.File) (*Device, error) {
	uniq, _ := getUniq(f)
	name, _ := getName(f)
	phys, _ := getPhys(f)
	gid, _ := getId(f)
	v, _ := getVersion(f)

	capMap := getCapMap(f)

	version := fmt.Sprintf("%d.%d.%d", v[0], v[1], v[2])

	return &Device{
		file:         f,
		Name:         name,
		Uniq:         uniq,
		Phys:         phys,
		InputId:      *gid,
		Version:      version,
		Capabilities: capMap,
	}, nil

}

func (d Device) String() string {
	return fmt.Sprintf(
		`
Name: %s
Phys: %s
Uniq: %s
Prop: %s
Drive Version: %s
File: %v
InputId: %v`, d.Name, d.Phys, d.Uniq, d.Prop, d.Version, d.file.Name(), d.InputId)

}
