package advsync

import (
	"sync"

	"github.com/puzpuzpuz/xsync/v4"
)

// NamedMutexSM provides an independent mutex for each key using xsync.Map.
type NamedMutexSM[K comparable] struct {
	internalMap *xsync.Map[K, *sync.Mutex]
}

// NewNamedMutexSM creates an empty keyed mutex collection.
func NewNamedMutexSM[K comparable]() NamedMutexSM[K] {
	return NamedMutexSM[K]{
		xsync.NewMap[K, *sync.Mutex](),
	}
}

// Unlock releases the mutex associated with slug.
//
// Parameters:
//   - slug: key identifying the mutex to release.
func (nm *NamedMutexSM[K]) Unlock(slug K) {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.Mutex{})
	mutex.Unlock()
}

// UnlockSafe releases the mutex associated with slug when it appears locked.
//
// Parameters:
//   - slug: key identifying the mutex to release.
//
// Returns:
//   - bool: true when the mutex was released; false when it was not locked.
func (nm *NamedMutexSM[K]) UnlockSafe(slug K) bool {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.Mutex{})
	return unlockSafe(mutex)
}

// Lock acquires the mutex associated with slug.
//
// Parameters:
//   - slug: key identifying the mutex to acquire.
func (nm *NamedMutexSM[K]) Lock(slug K) {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.Mutex{})
	mutex.Lock()
}
