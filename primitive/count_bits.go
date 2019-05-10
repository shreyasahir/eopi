package main

import "fmt"

func main() {
	x := 7
	fmt.Println("countbits", countBits(x))
}

func countBits(x int) int {
	nbits := 0
	for x != 0 {
		nbits += x & 1
		x >>= 1
	}
	return nbits
}
