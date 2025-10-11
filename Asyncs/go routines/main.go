package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println("Task", id, "step", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	go func() {
		wg.Wait()
		fmt.Println("All done")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Elapsed:", time.Since(start))
	time.Sleep(300 * time.Millisecond)
}

func count(label string, limit int, d time.Duration) {
	for i := 0; i < limit; i++ {
		fmt.Println(label, i)
		time.Sleep(d)
	}
}

func main2() {
	go count("fast", 5, 100*time.Millisecond)
	go count("medium", 3, 200*time.Millisecond)
	go count("slow", 2, 300*time.Millisecond)
	time.Sleep(1 * time.Second)
	fmt.Println("end main")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("exit ok")
}
