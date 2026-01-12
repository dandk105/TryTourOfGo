package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	// 結果 0, 0, false, ""
	// 数字は0、、bool型はfalse, 文字列は空文字
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

