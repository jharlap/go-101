package main

import (
	"fmt"
	"time"
)

func restAndSpeak(i int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d ", i)
}

func main() {
	for i := 1; i <= 5; i++ {
		restAndSpeak(i)
	}
	fmt.Println()
}
