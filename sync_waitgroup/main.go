package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("HELLO FROM WORKER", i)
	time.Sleep(2000 * time.Millisecond)
	fmt.Printf("Worker %d done \n", i)
}

func main() {
	fmt.Println("HELLO FROM SYNC WAITGROUP")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("ALL DONE")
}
