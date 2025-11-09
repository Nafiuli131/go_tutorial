package main

import "fmt"

// T is a type parameter — can be any type that supports + operator
func add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(5, 10))       // int
	fmt.Println(add(3.2, 4.8))    // float64
}
