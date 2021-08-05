package main

//go:noinline
func add(firstArg, secondArg int32) (int32, bool) {
	return firstArg + secondArg, true
}

func main() { add(10, 32) }
