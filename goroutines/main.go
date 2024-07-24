package main

import (
	"fmt"
	"time"
)

func pri() {
	fmt.Println("HELLO FROM PRI FUNC")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("HELLO AFTER TIME SLEEP FUNC")
}

func SecondPri() {
	fmt.Println("HELLO FROM SECOUND PRI FUNC")
}

func main() {
	fmt.Println("goroutines")
	go pri()
	SecondPri()

	time.Sleep(3000 * time.Millisecond)
}
