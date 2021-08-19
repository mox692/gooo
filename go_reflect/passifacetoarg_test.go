package go_reflect

import (
	"reflect"
	"testing"
)

// AIM: interfaceを引数に渡すと遅くなりそう？を確かめる
// 結果: ifaceを引数に渡すということが直接的にperformanceに影響することはない。
//      (ただ↓の例だとstackを確保してる(関数を呼んでる)ので、それで遅くなることhじゃあるかも)
type a struct {
	name string
	age  int
}

func passIface(v interface{}) {
	reflect.ValueOf(v).Elem()
}

func BenchmarkCallDirectly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := a{}
		reflect.ValueOf(&b).Elem()
	}
}

func BenchmarkCallIndirectly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := a{}
		passIface(&b)
	}
}
