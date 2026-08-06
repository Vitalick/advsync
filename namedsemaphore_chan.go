package advsync

import "sync"

// NamedSemaphoreChan provides an independent channel-based semaphore for each key.
type NamedSemaphoreChan[K comparable] struct {
	mapLock     sync.RWMutex
	maxCount    uint
	internalMap map[K]*SemaphoreChan
}

// NewNamedSemaphoreChan creates a keyed channel-based semaphore collection.
//
// Parameters:
//   - maxCount: maximum simultaneous acquisitions for each key.
func NewNamedSemaphoreChan[K comparable](maxCount uint) *NamedSemaphoreChan[K] {
	return &NamedSemaphoreChan[K]{
		internalMap: map[K]*SemaphoreChan{},
		maxCount:    maxCount,
	}
}

// Release releases one acquisition for the channel-based semaphore associated with slug.
//
// Parameters:
//   - slug: key identifying the semaphore to release.
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

// Acquire waits for and acquires a permit from the semaphore associated with slug.
//
// Parameters:
//   - slug: key identifying the semaphore to acquire.
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
