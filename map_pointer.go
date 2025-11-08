package main

import "fmt"

func update(m map[string]int) {
    m["A"] = 100
}

func main() {
    myMap := map[string]int{"A": 1, "B": 2}
    update(myMap)
    fmt.Println(myMap) // map[A:100 B:2]
}

