package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first := 0
	second := 1

	return func() int {
		temp := second
		second += first
		first = temp
		return second
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
> go run exercise-fibonacci-closure.go
1
2
3
5
8
13
21
34
55
89
*/