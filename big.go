// +build s390x ppc64 mips64 mips

package cpuendian

import "encoding/binary"

const Big = false
const Little = true

var CPU binary.BigEndian
