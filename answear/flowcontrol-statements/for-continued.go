package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; { // 初期化ステートメントを省略
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
