package building_blocks

import (
	"fmt"
	"sync"
)

// You can think of WaitGroup as concurrent-safe counter

func usingWaitGroup() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %d\n", id)
	}

	numCounter := 5
	var wg sync.WaitGroup
	wg.Add(numCounter)
	for i := 0; i < numCounter; i++ {
		go hello(&wg, i + 1)
	}

	wg.Wait()
}
