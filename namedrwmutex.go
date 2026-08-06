package advsync

import (
	"sync"
)

// NamedRWMutex provides an independent read/write mutex for each key.
type NamedRWMutex[K comparable] struct {
	mapLock     sync.RWMutex
	internalMap map[K]*sync.RWMutex
}

// NewNamedRWMutex creates an empty keyed read/write mutex collection.
func NewNamedRWMutex[K comparable]() *NamedRWMutex[K] {
	return &NamedRWMutex[K]{
		internalMap: map[K]*sync.RWMutex{},
	}
}

// Unlock releases the write lock associated with slug.
//
// Parameters:
//   - slug: key identifying the read/write mutex to release.
func (nm *NamedRWMutex[K]) Unlock(slug K) {
	nm.mapLock.RLock()
	mutex, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.RWMutex{}
		nm.internalMap[slug].Unlock()
		nm.mapLock.Unlock()
		return
	}
	mutex.Unlock()
}

// UnlockSafe releases the write lock associated with slug when it appears locked.
//
// Parameters:
//   - slug: key identifying the read/write mutex to release.
//
// Returns:
//   - bool: true when the write lock was released; false when it was not locked.
func (nm *NamedRWMutex[K]) UnlockSafe(slug K) bool {
	nm.mapLock.RLock()
	mutex, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		return false
	}
	return unlockSafeRW(mutex)
}

// Lock acquires the write lock associated with slug.
//
// Parameters:
//   - slug: key identifying the read/write mutex to acquire.
func (nm *NamedRWMutex[K]) Lock(slug K) {
	nm.mapLock.RLock()
	mutex, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.RWMutex{}
		nm.internalMap[slug].Lock()
		nm.mapLock.Unlock()
		return
	}
	mutex.Lock()
}

// RUnlock releases a read lock associated with slug.
//
// Parameters:
//   - slug: key identifying the read/write mutex to release.
func (nm *NamedRWMutex[K]) RUnlock(slug K) {
	nm.mapLock.RLock()
	mutex, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.RWMutex{}
		nm.internalMap[slug].RUnlock()
		nm.mapLock.Unlock()
		return
	}
	mutex.RUnlock()

}

// RUnlockSafe releases a read lock associated with slug when one is held.
//
// Parameters:
//   - slug: key identifying the read/write mutex to release.
//
// Returns:
//   - bool: true when a read lock was released; false when no reader was present.
func (nm *NamedRWMutex[K]) RUnlockSafe(slug K) bool {
	nm.mapLock.RLock()
	mutex, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		return false
	}
	return rUnlockSafeRW(mutex)
}

// RLock acquires a read lock associated with slug.
//
// Parameters:
//   - slug: key identifying the read/write mutex to acquire.
func (nm *NamedRWMutex[K]) RLock(slug K) {
	nm.mapLock.RLock()
	mutex, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.RWMutex{}
		nm.internalMap[slug].RLock()
		nm.mapLock.Unlock()
		return
	}
	mutex.RLock()

}
