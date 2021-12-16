package advsync

import "sync"

//NamedSemaphore is a named semaphore via sync.RWMutex
type NamedSemaphore struct {
	mapLock     sync.RWMutex
	maxCount    uint
	internalMap map[interface{}]*Semaphore
}

//NewNamedSemaphore create new named semaphore
func NewNamedSemaphore(maxCount uint) *NamedSemaphore {
	return &NamedSemaphore{
		internalMap: map[interface{}]*Semaphore{},
		maxCount:    maxCount,
	}
}

//Release semaphore by name
func (nm *NamedSemaphore) Release(slug interface{}) error {
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

//Acquire semaphore by name
func (nm *NamedSemaphore) Acquire(slug interface{}) {
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
