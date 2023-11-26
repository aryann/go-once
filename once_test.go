package once_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"aryan.app/go-once"
)

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

func TestSyncOnce(t *testing.T) {
	var once sync.Once
	run(t, &once, 100)
}

func TestMutexBasedOnce(t *testing.T) {
	var once once.MutexBasedOnce
	run(t, &once, 100)
}

var routineCounts = []int{1, 1e3, 1e5, 1e7}

type newOnce func() Once

func benchmarkOnce(b *testing.B, new newOnce) {
	b.Helper()
	for _, count := range routineCounts {
		b.Run(fmt.Sprintf("routine_count_%d", count), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				run(b, new(), count)
			}
		})
	}
}

func BenchmarkSyncOnce(b *testing.B) {
	benchmarkOnce(b, func() Once {
		return &sync.Once{}
	})
}

func BenchmarkMutexBasedOnce(b *testing.B) {
	benchmarkOnce(b, func() Once {
		return &once.MutexBasedOnce{}
	})
}
