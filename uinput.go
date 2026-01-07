package main

import (
	"encoding/binary"
	"os"
	"syscall"
	"unsafe"
)

/* ===== Linux uinput 常量（来自 linux/uinput.h + input-event-codes.h） ===== */

const (
	UI_SET_EVBIT  = 0x40045564
	UI_SET_KEYBIT = 0x40045565
	UI_DEV_CREATE = 0x5501

	EV_KEY = 0x01
	EV_SYN = 0x00

	SYN_REPORT = 0
)

/* ======================== */

type UInput struct {
	fd *os.File
}

func NewUInput() (*UInput, error) {
	fd, err := syscall.Open("/dev/uinput", syscall.O_WRONLY|syscall.O_NONBLOCK, 0)
	if err != nil {
		return nil, err
	}

	f := os.NewFile(uintptr(fd), "/dev/uinput")

	// Enable key events
	ioctl(f.Fd(), UI_SET_EVBIT, EV_KEY)

	// Enable all keys (0-255 足够键盘用了)
	for i := 0; i < 256; i++ {
		ioctl(f.Fd(), UI_SET_KEYBIT, uintptr(i))
	}

	var uidev uinputUserDev
	copy(uidev.Name[:], []byte("ayanokey"))
	uidev.ID.Bustype = 0x03 // BUS_USB
	uidev.ID.Vendor = 0x1234
	uidev.ID.Product = 0x5678
	uidev.ID.Version = 1

	if err := binary.Write(f, binary.LittleEndian, &uidev); err != nil {
		return nil, err
	}

	ioctl(f.Fd(), UI_DEV_CREATE, 0)

	return &UInput{fd: f}, nil
}

func (u *UInput) SendKey(code uint16, value int32) {
	ev := inputEvent{
		Type:  EV_KEY,
		Code:  code,
		Value: value,
	}
	binary.Write(u.fd, binary.LittleEndian, &ev)

	syn := inputEvent{
		Type:  EV_SYN,
		Code:  SYN_REPORT,
		Value: 0,
	}
	binary.Write(u.fd, binary.LittleEndian, &syn)
}

func ioctl(fd uintptr, cmd uintptr, arg uintptr) {
	syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, arg)
}

/* ===== structs ===== */

type inputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

type inputID struct {
	Bustype uint16
	Vendor  uint16
	Product uint16
	Version uint16
}

type uinputUserDev struct {
	Name         [80]byte
	ID           inputID
	FFEffectsMax uint32
	Absmax       [64]int32
	Absmin       [64]int32
	Absfuzz      [64]int32
	Absflat      [64]int32
}

/* 防止 go vet 抱怨 unused unsafe（uinput 常见问题） */
var _ = unsafe.Sizeof(0)

