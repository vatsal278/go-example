package main

import (
	"fmt"
	"log"
	"sync"
)

var y int64

func main() {
	var m sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i <= 20; i++ {

		wg.Add(1)
		//fmt.Println("goroutine ", i)

		go func(wg *sync.WaitGroup, m *sync.Mutex) {
			for c := 0; c < 5000; c++ {
				m.Lock()
				y = y + 1
				log.Print(y)
				m.Unlock()
			}
			wg.Done()
		}(&wg, &m)

	}

	wg.Wait()

	fmt.Println("ops:", y)
}
