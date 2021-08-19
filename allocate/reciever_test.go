package allocate

import "testing"

// AIM: レシーバを色々変えてみる
// 結果: ptrレシーバだろうが、値だろうが、関数内で宣言した変数のptrを返すとescapeされる。。
// 考察: local変数は基本的にはstackに置かれるが、そのアドレスが返される(stackのscope外に出る)となると
//       その値をどこか別の場所(heap)においておく必要が出てくる

type person struct {
	name string
	age  int8
	id   int
}

func (h person) getAge() int8   { return h.age }
func (h *person) getAgeP() int8 { return h.age }

func (h person) setAge(age int8)   { h.age = age }
func (h *person) setAgeP(age int8) { h.age = age }

func (h person) retPtr() *person {
	return &person{name: h.name, age: 33}
}

var p = person{}

func BenchmarkGetAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.getAge()
	}
}

func BenchmarkSetAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.setAge(3)
	}
}

func BenchmarkSetAgeP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.setAgeP(3)
	}
}

// このパターンはやはりescapeされる
func BenchmarkRetPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.retPtr()
	}
}
