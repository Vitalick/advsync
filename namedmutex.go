package advsync

import (
	"sync"
)

// NamedMutex is a named mutex via sync.RWMutex
type NamedMutex[K comparable] struct {
	mapLock     sync.RWMutex
	internalMap map[K]*sync.Mutex
}

// NewNamedMutex create new named mutex
func NewNamedMutex[K comparable]() *NamedMutex[K] {
	return &NamedMutex[K]{
		internalMap: map[K]*sync.Mutex{},
	}
}

// Unlock mutex by name
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

// UnlockSafe mutex by name
func (nm *NamedMutex[K]) UnlockSafe(slug K) bool {
	nm.mapLock.RLock()
	mutex, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		return false
	}
	return unlockSafe(mutex)
}

// Lock mutex by name
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
