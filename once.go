package once

import (
	"sync"
	"sync/atomic"
)

type MutexOnlyOnce struct {
	mutex sync.Mutex
	done  bool
}

func (m *MutexOnlyOnce) Do(f func()) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if !m.done {
		f()
		m.done = true
	}
}

type MutexAndInt32AtomicOnce struct {
	mutex sync.Mutex
	done  int32
}

func (m *MutexAndInt32AtomicOnce) Do(f func()) {
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

type MutexAndInt64AtomicOnce struct {
	mutex sync.Mutex
	done  int64
}

func (m *MutexAndInt64AtomicOnce) Do(f func()) {
	if atomic.LoadInt64(&m.done) == 1 {
		return
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	if atomic.LoadInt64(&m.done) == 1 {
		return
	}
	f()
	atomic.StoreInt64(&m.done, 1)
}
