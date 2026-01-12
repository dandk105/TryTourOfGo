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

# Basic types
Go言語の基本の型
int,uintとuintptrが基本32ビットだが、64ビットのシステムだと64ビットに拡張する
整数値が必要な時で特別な理由がない時は intを利用する様にする。
**ビットシフト計算** 

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

# Zero values
ゼロの値の定義は型によって違う
- 数字型は0
- 文字列は空文字""
- bool型はfalse

# Type conversions
Go言語は型を明示的に型変換する必要がある
暗黙的に別の型に変換する事はできない。
参考例

```
i := 42
f := i        // ❌ エラー！
u := f        // ❌ エラー！
```

# Constants
定数constantsは := syntaxを利用する事ができない。
必ず等式で入力する必要がある。

# Numeric Constants
大きい数値だと表現できる64bitの枠以上の値は定義は可能だが、コンパイルする際にエラーになる。
浮動小数点は逆に精度が落ちるがコンパイルしてもエラーにはならない。
