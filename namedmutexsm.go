package advsync

import (
	"github.com/puzpuzpuz/xsync/v3"
	"sync"
)

// NamedMutexSM is a named mutex via sync.Map
type NamedMutexSM[K comparable] struct {
	internalMap xsync.MapOf[K, *sync.Mutex]
}

// Unlock mutex by name
func (nm *NamedMutexSM[K]) Unlock(slug K) {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.Mutex{})
	mutex.Unlock()
}

// UnlockSafe mutex by name
func (nm *NamedMutexSM[K]) UnlockSafe(slug K) bool {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.Mutex{})
	return unlockSafe(mutex)
}

// Lock mutex by name
func (nm *NamedMutexSM[K]) Lock(slug K) {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.Mutex{})
	mutex.Lock()
}
