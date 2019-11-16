package main

import (
	"fmt"
)

func main() {
	c := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	d := make([]int, 5)
	for i := 0; i < 5; i++ {
		d = append(d, i)
		fmt.Println(d)
	}
	fmt.Println(d)
}
