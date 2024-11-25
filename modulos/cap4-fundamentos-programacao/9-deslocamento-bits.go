package main

import "fmt"

func main() {
	y := 24
	x := y << 2
	z := y >> 2
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", z)
}

// 0110000
// 0011000
// 0001100
// 0000110
// 0000011
