package main

import "fmt"

func main() {
	numGenerator := generator()

	for i := 0; i < 5; i++ {
		fmt.Print(numGenerator(), "\t")
	}
}

// this function return another function
func generator() func() int {
	var i = 0
	return func() int {
		i++
		return i
	}
}
