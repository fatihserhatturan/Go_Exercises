package main

import "fmt"

var globalVar = "I am global"

func localScopeExample() {
	localVar := "I am local"
	fmt.Println(localVar)
	{
		innerVar := "I am inner"
		fmt.Println(innerVar)
	}
}

func shadowExample() {
	value := 10
	fmt.Println("outer value:", value)
	{
		value := 20
		fmt.Println("inner value:", value)
	}
	fmt.Println("after block:", value)
}

func main() {
	fmt.Println(globalVar)
	localScopeExample()
	shadowExample()
	for i := 0; i < 3; i++ {
		fmt.Println("loop index:", i)
	}
	fmt.Println("done")
}
