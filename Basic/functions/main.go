package main

import "fmt"

func main() {

total:= add(10,20)
fmt.Println(total)

allTotal,difference:= calc(10,20)
fmt.Println(allTotal,difference)

}

func add(x int, y int) int {

	fmt.Println(x+y)

	return x+y

}

func calc(x int , y int) (int,int){

	return x+y,x-y
}
