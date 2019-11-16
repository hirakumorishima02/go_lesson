package main

import (
	"fmt"
)

func main() {
	// 配列を入れ子にするときには[][]と2つ書く
	// 代入された配列([]int{0,1,2}...など）は1つ目の[]に入る
	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)
	board = append(board, []int{9, 10, 11}, []int{12, 13, 14})
	fmt.Println(board)
}
