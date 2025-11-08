package main

import (
	"fmt"
)

func modify(slice []int){
	slice[0]=99
}

func main(){
	arr:=[]int{1,2,3}
	modify(arr)
	fmt.Print(arr)
}