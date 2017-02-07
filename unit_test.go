package cpuendian

import (
	"testing"
	"unsafe"
)

func TestBinary(t *testing.T) {
	if Big != !Little {
		t.Error("not either-or")
	}
}

func TestEndianness(t *testing.T) {
	x := uint32(0x12345678)
	raw := *(*[4]byte)(unsafe.Pointer(&x))
	switch raw {
	case [4]byte{0x12, 0x34, 0x56, 0x78}:
		if Little {
			t.Error("CPU is big-endian")
		}
	case [4]byte{0x78, 0x56, 0x34, 0x12}:
		if Big {
			t.Error("CPU is little-endian")
		}
	default:
		t.Error("non-trivial endianess", raw)
	}
}
