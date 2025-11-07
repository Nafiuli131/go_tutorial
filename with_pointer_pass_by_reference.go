package main

import "fmt"

func increment(n *int) {
    *n = *n + 1
}

func main() {
    x := 10
    increment(&x)
    fmt.Println(x) // 11
}

//Now we passed the address of x, so the function modified the original variable.
