package allocate

import (
	"testing"
	"unsafe"
)

type Foo struct {
	S *string
}

func (f *Foo) String() string {
	return *f.S
}

type FooTrick struct {
	S unsafe.Pointer
}

func (f *FooTrick) String() string {
	return *(*string)(f.S)
}

func NewFoo(s string) Foo {
	// local変数sのアドレスをFooに詰めて返してるので、sはheapへescape.
	return Foo{S: &s}
}

func NewFooTrick(s string) FooTrick {
	// これも↑と同じでheapにescapeされるはずだが....。
	return FooTrick{S: noescape(unsafe.Pointer(&s))}
}

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	// return unsafe.Pointer(x ^ 0)
	return unsafe.Pointer(x)
}

var s string = "hello"

func BenchmarkNewFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFoo(s)
	}
}

func BenchmarkNewFooTrick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFooTrick(s)
	}
}

// func main() {
// 	f1 := NewFoo(s)
// 	f2 := NewFooTrick(s)
// 	f1.String()
// 	f2.String()
// }
