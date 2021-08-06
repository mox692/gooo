package main

import (
	"fmt"
)

func add32(a int32) int32
func add32_2(a, b int32) int32
func add64(a int64) int64

// caller get int64 value and return it directory.(ret1)
// And also, caller call callee internaly,
// and return callee's result as ret2.
func caller(i int32) (ret1 int32, ret2 int32)
func callee(a int32) int32 {
	return -a
}

func caller2(i int64) (ret1 int64, ret2 int64)
func callee2(a int64) int64 {
	return a * a
}

func sliceBase(b []byte) uintptr
func sliceLen(b []byte) int64
func sliceCap(b []byte) int64

func main() {
	fmt.Println(add32(3))
	fmt.Println(add32_2(3, 3))
	fmt.Println(add64(3))

	s := make([]byte, 10, 20)
	fmt.Println(sliceLen([]byte(s)))
	fmt.Println(sliceCap([]byte(s)))
	fmt.Println(sliceBase([]byte(s)))

	p := Point{
		X: 3,
		Y: 6,
	}
	fmt.Println(addPoint(p))

	fmt.Println(caller(6))   // 8, -8
	fmt.Println(caller2(10)) // 10, 100
}
