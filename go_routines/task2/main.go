package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	var y int64
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 10; c++ {
				y = y + 1
				log.Print(y)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", y)
}
