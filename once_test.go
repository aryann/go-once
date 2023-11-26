package once_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"aryan.app/go-once"
)

type newOnce func() Once

var implementations = []struct {
	name string
	new  newOnce
}{
	{
		name: "standard-lib-once",
		new:  func() Once { return &sync.Once{} },
	},
	{
		name: "mutex-only-once",
		new:  func() Once { return &once.MutexOnlyOnce{} },
	},
	{
		name: "mutex-and-int32-atomic-once",
		new:  func() Once { return &once.MutexAndInt32AtomicOnce{} },
	},
	{
		name: "mutex-and-int64-atomic-once",
		new:  func() Once { return &once.MutexAndInt64AtomicOnce{} },
	},
}

type Once interface {
	Do(f func())
}

type TestHelper interface {
	Helper()
	Fatalf(format string, args ...any)
}

func run(test TestHelper, once Once, routineCount int) {
	test.Helper()

	var count int64

	var wg sync.WaitGroup
	for i := 0; i < routineCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(func() {
				atomic.AddInt64(&count, 1)
			})
		}()
	}

	wg.Wait()

	finalCount := atomic.LoadInt64(&count)
	if finalCount != 1 {
		test.Fatalf("Once must have been invoked exactly once, not %d time(s)", finalCount)
	}
}

func Test(t *testing.T) {
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			run(t, implementation.new(), 100)

		})
	}
}

var routineCounts = []int{1, 1e3, 1e5, 1e7}

func Benchmark(b *testing.B) {
	for _, implementation := range implementations {
		for _, count := range routineCounts {
			b.Run(fmt.Sprintf("%s/routines=%d/", implementation.name, count), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					run(b, implementation.new(), count)
				}
			})
		}
	}
}
