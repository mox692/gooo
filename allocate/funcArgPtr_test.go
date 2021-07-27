package allocate

import "testing"

// AIM: ptr渡しは本当にコストが高いのか
// 結果:
// 引数がprimitive型だった場合はそこまで速度に違いがない
// ただ、sizeの大きいstructなどだと、パフォーマンスに差が出た

type a struct {
	b int64
	c string
	f string
	g string
	h int64
	d []b
	e []string
}

type b struct {
	c string
	e []string
	f string
	g string
	h int64
}

func ArgVal(v a) {
}

func ArgPtr(v *a) {

}

func BenchmarkArgVal(b *testing.B) {
	aa := a{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ArgVal(aa)
	}
}

func BenchmarkArgPtr(b *testing.B) {
	aa := a{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ArgPtr(&aa)
	}
}
