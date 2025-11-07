package main

import ("fmt"
 "unsafe")

func main(){
	var x int =10
	var y int =12
	var p *int =&x
	var q *int
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:",&x)
	fmt.Println("Pointer p:", p)
    fmt.Println("Value via pointer:", *p)
	fmt.Println("Pointer q:", q)
	p=&y
	fmt.Println("modified value:",*p)
	fmt.Println("address of modified value:",p)

	fmt.Printf("Size of pointer to int: %d bytes\n", unsafe.Sizeof(p))
    fmt.Printf("Size of pointer to float64: %d bytes\n", unsafe.Sizeof(q))
}
