package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "a"
	names[1] = "b"
	names[2] = "c"

	fmt.Println(names)

	var slicerNames = []string{"a", "b", "c"}

	fmt.Println(slicerNames)

}
