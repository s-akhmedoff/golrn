package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var x, y int = 0, 1

	return func() int {
		x, y = y, x+y
		return x
	}
}

func Fib() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
