//go:build ppc64 || s390x || mips || mips64
// +build ppc64 s390x mips mips64

// note list of big-endian GOARCH comes from <go compiler>/src/syscall/endian_big.go

package cpuendian

import "encoding/binary"

const Big = false
const Little = true

var CPU binary.BigEndian
