package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffer size 2

	ch <- 10
	ch <- 20

	val1 := <-ch
	val2 := <-ch

	fmt.Println(val1, val2)

}
