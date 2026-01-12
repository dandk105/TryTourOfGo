package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	// 1<<64がなんで2の64乗になるんだ？
	// 2進数で左ビットシフトが肝だった。
	// 1    = 000000001 (2進数)　で
	// 1<<1 = 000000010 -> 2になる。これを64回繰り返すという意味なので2の64乗になる。
	MaxInt uint64     = 1<<64 - 1
	// 複素数の平方根を計算するライブラリ
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

