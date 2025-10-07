package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {

    p1 := Person{"Fatih", 25}

    p2 := Person{Name: "Ahmet", Age: 30}

    var p3 Person
    p3.Name = "Zeynep"
    p3.Age = 22

    fmt.Println(p1)
    fmt.Println(p2)
    fmt.Println(p3)

    fmt.Println("p1'in adı:", p1.Name)
    fmt.Println("p2'nin yaşı:", p2.Age)
}
