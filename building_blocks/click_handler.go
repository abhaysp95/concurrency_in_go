package building_blocks

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func ClickHandler() {
	button := Button {
		sync.NewCond(&sync.Mutex{}),
	}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {  // S1
		fmt.Println("Maximize Window")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {  // S2
		fmt.Println("Show Dialog Box")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {  // S#
		fmt.Println("Call Api")
		clickRegistered.Done()
	})
	button.Clicked.Broadcast()

	clickRegistered.Wait()
}
