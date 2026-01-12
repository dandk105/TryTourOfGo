package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// 型の変更は 変更先の型(変更したい値) の書き方で可能
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

