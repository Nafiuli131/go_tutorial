package main

import "fmt"

func main(){
	var testString string = "Hello"
	fmt.Println(testString)
	fmt.Println(testString[0])
	fmt.Println(testString[1])
	//print actual character
	fmt.Println(string(testString[0]))
}
