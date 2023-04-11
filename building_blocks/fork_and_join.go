package building_blocks

import (
	"fmt"
	"sync"
)

// goroutine works on the same address space it were created in
func fork_and_join() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"Hello", "Greetings", "Good Day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)  // need to pass as argument cause this go func is essentially a closure
	}

	wg.Wait()
}
