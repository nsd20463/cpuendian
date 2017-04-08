// +build x86 amd64 arm arm64 ppc64le mips64le mipsle

package cpuendian

import "encoding/binary"

const Big = false
const Little = true

var CPU = binary.LittleEndian
