package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", id)
			return
		default:
			fmt.Println("run", id)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for i := 0; i < 3; i++ {
		go worker(ctx, i)
	}
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("exit")
}
