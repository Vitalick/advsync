package advsync

import (
	"sync"
)

// NamedRWMutex is a named read/write mutex via sync.RWMutex
type NamedRWMutex[K comparable] struct {
	mapLock     sync.RWMutex
	internalMap map[K]*sync.RWMutex
}

// NewNamedRWMutex create new named read/write mutex
func NewNamedRWMutex[K comparable]() *NamedRWMutex[K] {
	return &NamedRWMutex[K]{
		internalMap: map[K]*sync.RWMutex{},
	}
}

// Unlock mutex by name
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

// UnlockSafe mutex by name
func (nm *NamedRWMutex[K]) UnlockSafe(slug K) bool {
	nm.mapLock.RLock()
	mutex, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		return false
	}
	return unlockSafeRW(mutex)
}

// Lock mutex by name
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

// RUnlock mutex by name
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

// RUnlockSafe mutex by name
func (nm *NamedRWMutex[K]) RUnlockSafe(slug K) bool {
	nm.mapLock.RLock()
	mutex, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		return false
	}
	return rUnlockSafeRW(mutex)
}

// RLock mutex by name
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
