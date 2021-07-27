package allocate

import "testing"

// go test -bench . -benchmem -gcflags "-m -l"で、両者にどれだけパフォーマンスの違いが出るか調べる

// escapeしてしまう
func RetPtr() *int {
	a := 4
	return &a
}

func RetVal() int {
	a := 4
	return a
}

func BenchmarkEscVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RetVal()
	}
}

func BenchmarkEscPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RetPtr()
	}
}
