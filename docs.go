/*
  Trivial package to define two compile-time constants
    Big
	Little
  one of which is true and the other false, depending on
  the endianness of the CPU, and a
    CPU     binary.ByteOrder
  with the CPU's endianness.

  This information is crucial in bit hacks, yet is lacking
  as a constant.

  Two constants is redundant, but having both makes the using code
  easier to read because it avoids negations in the if conditions.

  Copyright 2017 Nicolas S. Dade
*/
package cpuendian
