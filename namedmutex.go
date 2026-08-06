package advsync

import (
	"sync"
)

// NamedMutex provides an independent mutex for each key.
type NamedMutex[K comparable] struct {
	mapLock     sync.RWMutex
	internalMap map[K]*sync.Mutex
}

// NewNamedMutex creates an empty keyed mutex collection.
func NewNamedMutex[K comparable]() *NamedMutex[K] {
	return &NamedMutex[K]{
		internalMap: map[K]*sync.Mutex{},
	}
}

// Unlock releases the mutex associated with slug.
//
// Parameters:
//   - slug: key identifying the mutex to release.
func (nm *NamedMutex[K]) Unlock(slug K) {
	nm.mapLock.RLock()
	mutex, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.Mutex{}
		nm.internalMap[slug].Unlock()
		nm.mapLock.Unlock()
		return
	}
	mutex.Unlock()
}

// UnlockSafe releases the mutex associated with slug when it appears locked.
//
// Parameters:
//   - slug: key identifying the mutex to release.
//
// Returns:
//   - bool: true when the mutex was released; false when it was not locked.
func (nm *NamedMutex[K]) UnlockSafe(slug K) bool {
	nm.mapLock.RLock()
	mutex, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		return false
	}
	return unlockSafe(mutex)
}

// Lock acquires the mutex associated with slug.
//
// Parameters:
//   - slug: key identifying the mutex to acquire.
func (nm *NamedMutex[K]) Lock(slug K) {
	nm.mapLock.RLock()
	mutex, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.Mutex{}
		nm.internalMap[slug].Lock()
		nm.mapLock.Unlock()
		return
	}
	mutex.Lock()
}
