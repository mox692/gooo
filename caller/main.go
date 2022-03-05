package main

import (
	"fmt"
	"runtime"
)

// callerの情報をruntimeから引っ張ってくる.

func main() {
	pcs := make([]uintptr, 10)
	pcs = pcs[:runtime.Callers(0, pcs)]
	frames := runtime.CallersFrames(pcs)
	for {
		i, has := frames.Next()
		fmt.Printf("%+v\n", i)
		if !has {
			break
		}
	}
}
