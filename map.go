package main

import "fmt"

func main(){
	var myMapp map[string]int
	myMapp=make(map[string]int)
	myMapp["apple"]=5
	myMapp["banana"]=4
	for key,value := range myMapp {
		fmt.Println(key," ",value)
	}
}
