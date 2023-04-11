package building_blocks

import (
	"fmt"
	"sync"
)

// handling concurrency through memory access synchronization

func justMutexIt() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("decrementing: %d\n", count)
	}

	// increment
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	// decrement
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	wg.Wait()
}
