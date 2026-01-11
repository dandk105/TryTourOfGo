# 解いていて気づいた事

## Package
Exported namesでGoでは先頭の文字によって外部からアクセス可能か判断出来る
- 小文字でスタートしていると、そのパッケージのみでしか利用できない
- 大文字でスタートしていると、外部のパッケージからでも利用出来る

## Functions
2つ以上の連続した名前付き引数が**同じ型**を持つ場合、最後の1つを除いて型の記述を省略することができる
```golang

func sample(x int, y int)

↓

func sample(x, y int)
```

関数で2つの値を返却することが出来る
```golang

func sample(x int, y int)(string,string) {
    return y, x
}

```

名前つき戻り値をGoでは定義することが出来る。
メリットは戻り値の意味を明確に示すドキュメント代わりになる。


```golang
// どうなんだろ、なんかあんまり好きじゃない書き方。
// 先頭で戻り値を定義しているから、関数の中身を見なくても引数の概要を確認出来るというのは理解出来る。
// リファクタリングはしやすいかも？
// Warning: ただし、この書き方は短い関数でのみ利用可能。
// Product Codeでこの書き方されているのかな
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}


```

# Variables with initializers
右辺に初期化子がある場合は左辺の型宣言を省略することが出来る
```golang
	var c, python, java = true, false, "no!"
```
