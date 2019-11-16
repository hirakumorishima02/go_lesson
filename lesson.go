package main

import (
	"fmt"
)

func main() {
	var s [2]int
	s[0] = 100
	s[1] = 200
	fmt.Println(s)

	var n []int = []int{100, 200}
	n = append(n, 300)
	fmt.Println(n)
}
