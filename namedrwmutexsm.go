package advsync

import (
	"github.com/puzpuzpuz/xsync/v3"
	"sync"
)

// NamedRWMutexSM is a named read/write mutex via sync.Map
type NamedRWMutexSM[K comparable] struct {
	internalMap xsync.MapOf[K, *sync.RWMutex]
}

// Unlock mutex by name
func (nm *NamedRWMutexSM[K]) Unlock(slug K) {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	mutex.Unlock()
}

// UnlockSafe mutex by name
func (nm *NamedRWMutexSM[K]) UnlockSafe(slug K) bool {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	return unlockSafeRW(mutex)
}

// Lock mutex by name
func (nm *NamedRWMutexSM[K]) Lock(slug K) {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	mutex.Lock()
}

// RUnlock mutex by name
func (nm *NamedRWMutexSM[K]) RUnlock(slug K) {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	mutex.RUnlock()
}

// RUnlockSafe mutex by name
func (nm *NamedRWMutexSM[K]) RUnlockSafe(slug K) bool {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	return rUnlockSafeRW(mutex)
}

// RLock mutex by name
func (nm *NamedRWMutexSM[K]) RLock(slug K) {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	mutex.RLock()
}
