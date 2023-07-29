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

// for each once, if it's called once, don't call it again (interally uses an atomic counter)
func syncs_only_once() {
	var count int

	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Println("count: ", count)  // will print 1
}

func sync_once_deadlock() {
	var onceA sync.Once
	var onceB sync.Once

	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }

	onceA.Do(initA)  // deadlock
}

// sync.Once is not intended as guard for multiple initialization, but it
// guarantees that your functions are only called once
