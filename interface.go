package main

import "fmt"

type Greeter interface{
	Greet() string
}

type Person struct{
	Name string
}

func (p Person) Greet() string{
	return "Hello from human "+p.Name
}

type Dog struct {
	Name string
}

func (d Dog) Greet() string{
	return "Hello from dog "+d.Name
}

func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}

func main(){
	p1:=Person{"Nafiul"}
	SayHello(p1)
	d1:=Dog{"Tommy"}
	SayHello(d1)
}