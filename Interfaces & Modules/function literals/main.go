package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}

	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Println("Add:", add(3, 4))
	fmt.Println("Multiply:", multiply(2, 5))

	numbers := []int{1, 2, 3, 4, 5}
	total := 0
	for _, n := range numbers {
		func(x int) {
			total += x
		}(n)
	}

	fmt.Println("Total:", total)

	message := func(name string) string {
		return "Hello, " + name
	}
	fmt.Println(message("Gopher"))
}
