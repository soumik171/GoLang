package mutex

import (
	"fmt"
	"sync"
)

// Multiple goroutines access shared data at once, causing inconsistent results

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(w *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		w.Done()
	}()

	// Just lock the terms, that's does the modification

	p.mu.Lock()
	p.views++
}

func main() {

	var wg sync.WaitGroup
	myPost := post{
		views: 0,
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views)

}
