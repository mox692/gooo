package asm

type Mather interface {
    Add(a, b int32) int32
    Sub(a, b int64) int64
}

type Adder struct{ id int32 }
//go:noinline
func (adder Adder) Add(a, b int32) int32 { return a + b }
//go:noinline
func (adder Adder) Sub(a, b int64) int64 { return a - b }

func main() {
    m := Mather(Adder{id: 6754})

    // この関数呼び出しはインターフェースが実際にどのように利用されるかを確認するためだけのもの
    // この関数呼び出しがないと、リンカは Matherインターフェース が決して使われていないということを検知して実行ファイルに含めないように最適化してしまう
    m.Add(10, 32)
}