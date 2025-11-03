package main
import "fmt"

func main(){
	var array [3]int=[3]int{1,2,3}
	var i int
	for i=0;i<len(array)-1;i++{
		fmt.Print(i," ")
	}
}
