package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Greeter struct {
	Name string
}

func (g Greeter) Speak() string {
	return "Hello, " + g.Name
}

type Fareweller struct {
	Name string
}

func (f *Fareweller) Speak() string {
	return "Goodbye, " + f.Name
}

func Announce(s Speaker) {
	fmt.Println(s.Speak())
}

func Broadcast(list []Speaker) {
	for _, s := range list {
		switch v := s.(type) {
		case Greeter:
			fmt.Println("Greeter said:", v.Speak())
		case *Fareweller:
			fmt.Println("Fareweller said:", v.Speak())
		default:
			fmt.Println("Unknown speaker:", s.Speak())
		}
	}
}

func main() {
	g := Greeter{Name: "Alice"}
	f := &Fareweller{Name: "Bob"}
	Announce(g)
	Announce(f)
	people := []Speaker{g, f}
	Broadcast(people)
}
