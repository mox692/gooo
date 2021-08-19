package intendescape

import (
	"testing"
	_ "unsafe"
)

// AIM: reflect.escapesで、どうしてescapeが実現できるのかを知る.
//      さらに、どうしてpointerで渡した時と値で渡した時で挙動が異なるのかを知る.

//go:linkname escapesToHeap reflect.escapes
func escapesToHeap(v interface{})
func f(v interface{}) {}

func BenchmarkEscapeToHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 3
		escapesToHeap(&a)
	}
}

func BenchmarkNotEscapeToHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 3
		f(&a)
	}
}

/*
func escapes(x interface{}) {
	if dummy.b {
		dummy.x = x
	}
}

var dummy struct {
	b bool
	x interface{}
}
*/
