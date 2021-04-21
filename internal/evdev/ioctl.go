package evdev

import (
	"syscall"
	"unsafe"
)

const (
	iocNone  = 0x0
	iocWrite = 0x1
	iocRead  = 0x2

	iocTypeBits = 8
	iocNrShift  = 0
	iocSizeBits = 14
	iocNrBits   = 8
	iocDirBits  = 2

	iocTypeShift = iocNrShift + iocNrBits
	iocSizeShift = iocTypeShift + iocTypeBits
	iocDirShift  = iocSizeShift + iocSizeBits
	iocDirMask   = (1 << iocDirBits) - 1
)

func ioctl(fd uintptr, code uint32, ptr unsafe.Pointer) syscall.Errno {
	_, _, errNo := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(code), uintptr(ptr))
	return errNo
}

func ioc(dir, tp, nr, size uintptr) uint32 {
	var code uint32

	code |= uint32(dir) << 30
	code |= uint32(size) << 16
	code |= uint32(tp) << 8
	code |= uint32(nr)

	return code
}

func ior(tp, nr, tsize uintptr) uint32 {
	return ioc(iocRead, tp, nr, tsize)
}
