package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	count := 0

	for r := 0; r < 50; r++ {
		go func() {
			mutex.Lock()
			count += 1
			mutex.Unlock()
		}()
	}

	time.Sleep(time.Millisecond)
	fmt.Println("the count is : ", count)
}
