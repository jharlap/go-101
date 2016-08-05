package main

import (
	"fmt"
	"sync"
	"time"
)

func restAndSpeak(i int, wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d ", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go restAndSpeak(i, &wg)
	}
	wg.Wait()
	fmt.Println()
}
