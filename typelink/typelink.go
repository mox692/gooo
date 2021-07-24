package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//go:linkname typelinks reflect.typelinks
func typelinks() ([]unsafe.Pointer, [][]int32)

func main() {
	sections, offsets := typelinks()
	for i, base := range sections {
		for _, offset := range offsets[i] {
			typeAddr := uintptr(base) + uintptr(offset)
			typ := reflect.TypeOf(*(*interface{})(unsafe.Pointer(&typeAddr)))
			fmt.Println(typ)
		}
	}
}
