package allocate

import (
	"testing"
	"time"
)

// 関数が返すstructに着目.
// 単純にptrを返すようなcaseに加えて、Ptrフィールドをもつstructを返すようにしていると
// パフォーマンスが低下する.
// 作成した関数が意図せずptrをfieidに含むstructを返している場合は気付きにくいかも(time.Timeとか)

type item1 struct {
	name    string
	id      int
	getTime time.Time
}

type item2 struct {
	name    string
	id      int
	getTime int64
}

func innerFuncPtr() item1 {
	i := item1{}
	return i
}
func innerFuncVal() item2 {
	i := item2{}
	return i
}

func BenchmarkInnerFuncPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		innerFuncPtr()
	}
}
func BenchmarkInnerFuncVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		innerFuncVal()
	}
}
