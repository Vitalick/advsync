package advsync

import "sync"

// NamedSemaphore is a named semaphore via sync.RWMutex
type NamedSemaphore[K comparable] struct {
	mapLock     sync.RWMutex
	maxCount    uint
	internalMap map[K]*Semaphore
}

// NewNamedSemaphore create new named semaphore
func NewNamedSemaphore[K comparable](maxCount uint) *NamedSemaphore[K] {
	return &NamedSemaphore[K]{
		internalMap: map[K]*Semaphore{},
		maxCount:    maxCount,
	}
}

// Release semaphore by name
func (nm *NamedSemaphore[K]) Release(slug K) error {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = NewSemaphore(nm.maxCount)
		err := nm.internalMap[slug].Release()
		nm.mapLock.Unlock()
		return err
	}
	return locker.Release()
}

// Acquire semaphore by name
func (nm *NamedSemaphore[K]) Acquire(slug K) {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = NewSemaphore(nm.maxCount)
		nm.internalMap[slug].Acquire()
		nm.mapLock.Unlock()
		return
	}
	locker.Acquire()
}
