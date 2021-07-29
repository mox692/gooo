package allocate

import "testing"

// AIM: Interface型のmethod呼び出しは、type型のMethod呼び出しに比較して
// かなりコストが高いことを確認する.

type animal interface {
	Eat()
}

type human struct {
	name  string
	name1 string
	name2 string
	name3 string
	age   int
}

func (h human) Eat() {}

func callTypeMethod(h human) {
	h.Eat()
}

func callInterfaceMethod(a animal) {
	a.Eat()
}

var h = human{}

func BenchmarkCallTypeMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callTypeMethod(h)
	}
}

func BenchmarkCallInterfaceMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callInterfaceMethod(h) // escapes to heap
	}
}
