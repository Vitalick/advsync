package advsync

import "sync"

// NamedSemaphoreChan is a named semaphore via sync.RWMutex
type NamedSemaphoreChan struct {
	mapLock     sync.RWMutex
	maxCount    uint
	internalMap map[interface{}]*SemaphoreChan
}

// NewNamedSemaphoreChan create new named semaphore
func NewNamedSemaphoreChan(maxCount uint) *NamedSemaphoreChan {
	return &NamedSemaphoreChan{
		internalMap: map[interface{}]*SemaphoreChan{},
		maxCount:    maxCount,
	}
}

// Release semaphore by name
func (nm *NamedSemaphoreChan) Release(slug interface{}) {
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
func (nm *NamedSemaphoreChan) Acquire(slug interface{}) {
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
