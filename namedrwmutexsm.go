package advsync

import (
	"sync"

	"github.com/puzpuzpuz/xsync/v4"
)

// NamedRWMutexSM provides an independent read/write mutex for each key using xsync.Map.
type NamedRWMutexSM[K comparable] struct {
	internalMap *xsync.Map[K, *sync.RWMutex]
}

// NewNamedRWMutexSM creates an empty keyed read/write mutex collection.
func NewNamedRWMutexSM[K comparable]() NamedRWMutexSM[K] {
	return NamedRWMutexSM[K]{
		xsync.NewMap[K, *sync.RWMutex](),
	}
}

// Unlock releases the write lock associated with slug.
//
// Parameters:
//   - slug: key identifying the read/write mutex to release.
func (nm *NamedRWMutexSM[K]) Unlock(slug K) {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	mutex.Unlock()
}

// UnlockSafe releases the write lock associated with slug when it appears locked.
//
// Parameters:
//   - slug: key identifying the read/write mutex to release.
//
// Returns:
//   - bool: true when the write lock was released; false when it was not locked.
func (nm *NamedRWMutexSM[K]) UnlockSafe(slug K) bool {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	return unlockSafeRW(mutex)
}

// Lock acquires the write lock associated with slug.
//
// Parameters:
//   - slug: key identifying the read/write mutex to acquire.
func (nm *NamedRWMutexSM[K]) Lock(slug K) {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	mutex.Lock()
}

// RUnlock releases a read lock associated with slug.
//
// Parameters:
//   - slug: key identifying the read/write mutex to release.
func (nm *NamedRWMutexSM[K]) RUnlock(slug K) {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	mutex.RUnlock()
}

// RUnlockSafe releases a read lock associated with slug when one is held.
//
// Parameters:
//   - slug: key identifying the read/write mutex to release.
//
// Returns:
//   - bool: true when a read lock was released; false when no reader was present.
func (nm *NamedRWMutexSM[K]) RUnlockSafe(slug K) bool {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	return rUnlockSafeRW(mutex)
}

// RLock acquires a read lock associated with slug.
//
// Parameters:
//   - slug: key identifying the read/write mutex to acquire.
func (nm *NamedRWMutexSM[K]) RLock(slug K) {
	mutex, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	mutex.RLock()
}
