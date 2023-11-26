package once_test

import (
	"sync"
	"testing"
)

type Once interface {
	Do(f func())
}

func testOnce(t *testing.T, once Once) {
	t.Helper()

	done := make(chan struct{})
	defer close(done)

	count := make(chan struct{})
	invocations := 0

	go func() {
		for {
			_, ok := <-count
			if !ok {
				done <- struct{}{}
				return
			}
			invocations++
		}
	}()

	const numIterations = 100
	var wg sync.WaitGroup

	for i := 0; i < numIterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(func() {
				count <- struct{}{}
			})
		}()
	}

	wg.Wait()
	if invocations != 1 {
		t.Fatalf("Once must have been invoked exactly once, not %d time(s)", invocations)
	}

	close(count)
	<-done
}

func TestSyncOnce(t *testing.T) {
	var once sync.Once
	testOnce(t, &once)
}
