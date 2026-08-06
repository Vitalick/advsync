package advsync

import "sync"

// NamedSemaphore provides an independent Semaphore for each key.
type NamedSemaphore[K comparable] struct {
	mapLock     sync.RWMutex
	maxCount    uint
	internalMap map[K]*Semaphore
}

// NewNamedSemaphore creates a keyed semaphore collection with maxCount permits per key.
//
// Parameters:
//   - maxCount: maximum simultaneous acquisitions for each key.
func NewNamedSemaphore[K comparable](maxCount uint) *NamedSemaphore[K] {
	return &NamedSemaphore[K]{
		internalMap: map[K]*Semaphore{},
		maxCount:    maxCount,
	}
}

// Release releases one acquisition for the semaphore associated with slug.
//
// Parameters:
//   - slug: key identifying the semaphore to release.
//
// Returns:
//   - error: an error when no matching acquisition exists.
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

// Acquire waits for and acquires a permit from the semaphore associated with slug.
//
// Parameters:
//   - slug: key identifying the semaphore to acquire.
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
