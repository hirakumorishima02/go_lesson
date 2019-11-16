package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "14"
	// Stringからintに直す際にはAtoi関数を使用
	// Atoiは返り値に n と nil(エラー値)を返す
	// 第２引数には _ を入れる。errorなどを入れてしまうと「errorが使用されていない」というエラーが出るため
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v", i, i)
}
