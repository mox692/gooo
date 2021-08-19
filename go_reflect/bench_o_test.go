package go_reflect

import (
	"reflect"
	"testing"

	go_ref "github.com/goccy/go-reflect"
)

type bb struct {
	name string
	age  int
}

func valueFromReflectO(v interface{}) {
	reflect.ValueOf(v).Elem()
}

func valueFromGoReflectO(v interface{}) {
	go_ref.ValueNoEscapeOf(v).Elem()
}

func BenchmarkOGoPkg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := bb{}
		valueFromReflectO(&a)
	}
}

func BenchmarkOLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := bb{}
		valueFromGoReflectO(&a)
	}
}
