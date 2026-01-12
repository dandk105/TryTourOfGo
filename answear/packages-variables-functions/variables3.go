package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// ここの書き方は関数内でしか利用できない
	// 関数外だとエラーになる
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

