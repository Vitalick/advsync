package advsync

import "sync"

// NamedSemaphoreChan is a named semaphore via sync.RWMutex
type NamedSemaphoreChan[K comparable] struct {
	mapLock     sync.RWMutex
	maxCount    uint
	internalMap map[K]*SemaphoreChan
}

// NewNamedSemaphoreChan create new named semaphore
func NewNamedSemaphoreChan[K comparable](maxCount uint) *NamedSemaphoreChan[K] {
	return &NamedSemaphoreChan[K]{
		internalMap: map[K]*SemaphoreChan{},
		maxCount:    maxCount,
	}
}

// Release semaphore by name
func (nm *NamedSemaphoreChan[K]) Release(slug K) {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = NewSemaphoreChan(nm.maxCount)
		nm.internalMap[slug].Release()
		nm.mapLock.Unlock()
		return
	}
	locker.Release()
}

// Acquire semaphore by name
func (nm *NamedSemaphoreChan[K]) Acquire(slug K) {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = NewSemaphoreChan(nm.maxCount)
		nm.internalMap[slug].Acquire()
		nm.mapLock.Unlock()
		return
	}
	locker.Acquire()
}
