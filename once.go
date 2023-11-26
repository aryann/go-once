package once

import (
	"sync"
	"sync/atomic"
)

type MutexBasedOnce struct {
	mutex sync.Mutex
	done  bool
}

func (m *MutexBasedOnce) Do(f func()) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if !m.done {
		f()
		m.done = true
	}
}

type MutexBasedOnceWithInt32Atomic struct {
	mutex sync.Mutex
	done  int32
}

func (m *MutexBasedOnceWithInt32Atomic) Do(f func()) {
	if atomic.LoadInt32(&m.done) == 1 {
		return
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	if atomic.LoadInt32(&m.done) == 1 {
		return
	}
	f()
	atomic.StoreInt32(&m.done, 1)
}
