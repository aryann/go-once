package once

import "sync"

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
