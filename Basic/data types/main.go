package main

import "fmt"

func main() {
	var names map[string]int

	names = make(map[string]int, 0)

	names["a"] = 1
	names["b"] = 2
	names["c"] = 3

	fmt.Println(names["a"])

	names2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(names2)

}
