package main

import "fmt"

// user of closure
func fibonacci() func() int {
	b, c := 0, 1
	return func() int {
		b, c = c, b+c
		return c - b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
