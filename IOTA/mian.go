package main

import (
	"fmt"
)

// IOTA will increase by 1 so c1 is 0 and c2 is 1 ...
func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	// This wil do the same as above bbut automatic
	const (
		c12 = iota
		c22
		c32
	)
	fmt.Println(c12, c22, c32)

	// if you want to step by 2 or more you can do this
	const (
		c13 = iota * 2
		c23
		c33
	)
	fmt.Println(c13, c23, c33)

	// this will manule intarupt iota to start from 5 so the output will be 0 5 10 if you use 6 the output will be  0 6 12
	const (
		c14 = iota * 2
		c24 = iota * 5
		c35
	)
	fmt.Println(c14, c24, c35)

	// use can use negtive numbers

	const (
		y = -(iota + 2) // -2
		_               // we skip one
		x               // so this will be -4
	)
	fmt.Println(y, x)
}
