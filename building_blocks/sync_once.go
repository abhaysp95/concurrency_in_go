package building_blocks

import (
	"fmt"
	"sync"
)

func sync_once() {
	count := 0
	var once sync.Once

	increment := func() {
		count++
	}

	var increments sync.WaitGroup
	increments.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()

	fmt.Println("count: ", count)
}
