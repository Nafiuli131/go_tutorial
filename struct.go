package main

import "fmt"

type Person struct{
	Name string
	Age int
	Salary int
}

//You can attach functions to structs. These are called methods.
func (p Person) Greet() Person{
	return p
}

func main(){
	person1:=Person{"Nafiul",28,20}
	fmt.Println(person1.Age)
	person2:=Person{Name: "Nafi"}
	fmt.Println(person2.Greet())
}
