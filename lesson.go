package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "Hello World"
	fmt.Println(strings.Replace(s, "o", "X", 2))

	fmt.Println(`Test
	test
						test`)
	fmt.Println("\"")
	fmt.Println(`"`)
}
