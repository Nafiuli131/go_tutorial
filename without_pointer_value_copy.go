package main

import "fmt"
func increment(n int) {
    n = n + 1
	fmt.Println(n)
}

func main() {
    x := 10
    increment(x)
    fmt.Println(x) // still 10(so value is not changing)
}

//here Go passes a copy of x, not the original. So x never changes.