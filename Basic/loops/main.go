package main

import "fmt"

func main() {

	index := 1
	for index <= 10 {
		fmt.Println(index)
		index = index +1
	}

	total:=0

	for index:=1;index<=10;index++ {
		fmt.Println(index)
		total = total + index
	}

	fmt.Println(total)

	var numbers = []int{1,2,3,4}

	for _,value:= range numbers {
		fmt.Println(value)
	}


}
