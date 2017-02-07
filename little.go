// +build x86 amd64 arm arm64 ppc64le mips64le mipsle

/*
  Tiny package to define two compile-time constants
    Big
	Little
  one of which is true and the other false, depending on
  what the endiannesss of the CPU is.

  This information is crucial in bit hacks, yet is lacking
  as a constant.

  Two constants is redundant, but having both makes the using code
  easier to read because it avoids negations in the if conditions.

  Copyright 2017 Nicolas S. Dade
*/
package cpuendian

const Big = false
const Little = true
