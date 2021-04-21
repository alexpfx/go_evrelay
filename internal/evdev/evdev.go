package evdev

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

//https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h

type InputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

type InputMask struct {
	Type     uint32
	CodeSize uint32
	CodePtr  *uint64
}
type InputId struct {
	BusType uint16
	Vendor  uint16
	Product uint16
	Version uint16
}

func getVersion(file *os.File) ([]int, error) {
	out := int32(0)
	c := ior('E', 0x01, unsafe.Sizeof(out))
	err := ioctl(file.Fd(), c, unsafe.Pointer(&out))
	if err != 0 {
		fmt.Println(err)
		log.Fatalf("erro irrecuperável")
	}

	return extractVersion(out), err
}

func getId(file *os.File) (*InputId, error) {
	out := new(InputId)
	c := ior('E', 0x02, unsafe.Sizeof(*out))
	err := ioctl(file.Fd(), c, unsafe.Pointer(out))
	return out, err
}

func getName(file *os.File) (string, error) {
	out := new([256]byte)
	c := ior('E', 0x06, unsafe.Sizeof(*out))
	err := ioctl(file.Fd(), c, unsafe.Pointer(out))
	return string(out[:]), err
}

func getPhys(file *os.File) (string, error) {
	out := new([256]byte)
	c := ior('E', 0x07, unsafe.Sizeof(*out))
	err := ioctl(file.Fd(), c, unsafe.Pointer(out))
	return string(out[:]), err
}

func getUniq(file *os.File) (string, error) {
	out := new([256]byte)
	c := ior('E', 0x08, unsafe.Sizeof(*out))
	err := ioctl(file.Fd(), c, unsafe.Pointer(out))

	return string(out[:]), err
}

func getProp(file *os.File) ([]byte, error) {
	out := [256]byte{}
	c := ior('E', 0x09, unsafe.Sizeof(out))
	err := ioctl(file.Fd(), c, unsafe.Pointer(&out))
	return out[:], err
}
func getBit(file *os.File, evType uintptr) ([]byte, error) {
	out := [KeyMax]byte{}

	c := ior('E', 0x20+evType, unsafe.Sizeof(out))
	err := ioctl(file.Fd(), c, unsafe.Pointer(&out))

	return out[:], err
}

func getCapMap(f *os.File) map[uint16][]uint16 {
	capMap := make(map[uint16][]uint16, 0)

	for _, e := range eventTypes {
		caps, _ := getCapabilities(f, e)
		if len(caps) != 0 {
			capMap[e] = caps
		}
	}
	return capMap
}
func getCapabilities(file *os.File, evType uint16) ([]uint16, error) {
	evBits, _ := getBit(file, uintptr(evType))
	res := make([]uint16, 0)

	for i, n := range evBits {
		for j := 1; j <= 8; j++ {
			if isBitSet(byte(j), n) {
				res = append(res, uint16(i*8+j))
			}
		}
	}
	return res, nil
}

func (i InputId) String() string {
	return fmt.Sprintf(`
 BusType: 0x%x 
 Vendor: 0x%x 
 Product: 0x%x  
 Version: 0x%x`, i.BusType, i.Vendor, i.Product, i.Version)
}
