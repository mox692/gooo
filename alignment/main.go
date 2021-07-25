package main

// structのアライメントを、unsafe.Sizeof(), unsafe.AlignOf()等を用いて調べていく
import (
	"fmt"
	"unsafe"
)

type Foo struct {
	aaa bool  // 1
	bbb int32 // 4 (max)
	ссс bool  // 1
	ddd bool  // 1
} // sizeof Foo: 12

type Foo2 struct {
	aaa bool  // 1
	ссс bool  // 1
	ddd bool  // 1
	bbb int32 // 4 (max)
} // sizeof Foo2: 8

type Bar struct {
	aaa [2]bool // offset of bytes: 0
	ccc [2]bool // offset of bytes: 2
	bbb int32   // offset of bytes: 4
}

func main() {
	ff := Foo{}
	fff := Foo2{}
	bb := Bar{}
	fmt.Printf("sizeof Foo is %d\n", unsafe.Sizeof(ff))
	fmt.Printf("sizeof Foo2 is %d\n", unsafe.Sizeof(fff))
	fmt.Printf("sizeof Bar is %d\n", unsafe.Sizeof(bb))
	fmt.Printf("offsets of fields: aaa: %+v; bbb: %+v; ccc: %+v; ddd: %+v\n", unsafe.Offsetof(ff.aaa), unsafe.Offsetof(ff.bbb), unsafe.Offsetof(ff.ссс), unsafe.Offsetof(ff.ddd))
	fmt.Printf("offsets of fields: aaa: %+v; bbb: %+v; ccc: %+v; ddd: %+v\n", unsafe.Offsetof(fff.aaa), unsafe.Offsetof(fff.bbb), unsafe.Offsetof(fff.ссс), unsafe.Offsetof(fff.ddd))
	fmt.Printf("offsets of fields: aaa: %+v; ccc: %+v; bbb: %+v\n", unsafe.Offsetof(bb.aaa), unsafe.Offsetof(bb.ccc), unsafe.Offsetof(bb.bbb))
}
