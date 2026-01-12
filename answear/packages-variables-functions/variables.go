package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	// 実行結果: 0, false, false, false
	// MEMO: この実行結果から型を定義しない時の変数は自動的にbool型になる
	fmt.Println(i, c, python, java)
}

