//go:build 386 || amd64 || arm || arm64 || ppc64le || mips64le || mipsle || riscv64 || wasm
// +build 386 amd64 arm arm64 ppc64le mips64le mipsle riscv64 wasm

// note list of little-endian GOARCH comes from <go compiler>/src/syscall/endian_little.go
// note 2: the little-endian architecture 'amd64p32' was removed from the compiler in go 1.14.

package cpuendian

import "encoding/binary"

const Big = false
const Little = true

var CPU = binary.LittleEndian
