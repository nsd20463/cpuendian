//go:build 386 || amd64 || amd64p32 || arm || arm64 || ppc64le || mips64le || mipsle || wasm
// +build 386 amd64 amd64p32 arm arm64 ppc64le mips64le mipsle wasm

// note list of little-endian GOARCH comes from <go compiler>/src/syscall/endian_little.go

package cpuendian

import "encoding/binary"

const Big = false
const Little = true

var CPU = binary.LittleEndian
