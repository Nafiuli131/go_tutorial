package main

import "fmt"

func main(){
	var oddEvenFinder string = oddEven(431)
	fmt.Print(oddEvenFinder)
}

func oddEven(value int) string{
	if(value%2==0){
		return "It is an even number"
	}else{
		return "It is an odd number"
	}
}
