package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func main() {
	ch := make(chan int, 3)
	go producer(ch)
	timeout := time.After(1 * time.Second)
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("done")
				return
			}
			fmt.Println("recv", val)
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
}
