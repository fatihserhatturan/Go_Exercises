package main

import (
	"fmt"
	"time"
)

func sender(ch chan string) {
	values := []string{"x", "y", "z"}
	for _, v := range values {
		ch <- v
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func main() {
	ch := make(chan string)
	go sender(ch)
	for val := range ch {
		fmt.Println("got", val)
	}
	fmt.Println("closed")
	time.Sleep(100 * time.Millisecond)
}
