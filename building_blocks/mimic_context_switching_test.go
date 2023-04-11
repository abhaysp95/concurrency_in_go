package building_blocks

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitching(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	ch := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			ch <- token
		}
	}

	reciever := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-ch
		}
	}

	wg.Add(2)
	go sender()
	go reciever()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
