package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func changeAge(p *Person) {
    p.Age = 30
}

func main() {
    person := Person{Name: "Nafiul", Age: 25}
    changeAge(&person)
    fmt.Println(person.Age) // 30
}
