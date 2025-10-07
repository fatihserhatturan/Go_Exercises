package main

import (
	"fmt"
	"reflect"
)

func main() {

	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Go"
	quantity = 10
	discount = 0.37
	isInStock = true


	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity))
	fmt.Println(discount, reflect.TypeOf(discount))
	fmt.Println(isInStock, reflect.TypeOf(isInStock))



}
