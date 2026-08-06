package advsync

import (
	"github.com/puzpuzpuz/xsync/v4"
)

// NamedSemaphoreSM provides an independent Semaphore for each key using xsync.Map.
type NamedSemaphoreSM[K comparable] struct {
	internalMap *xsync.Map[K, *Semaphore]
	maxCount    uint
}

// NewNamedSemaphoreSM creates a keyed semaphore collection with maxCount permits per key.
//
// Parameters:
//   - maxCount: maximum simultaneous acquisitions for each key.
func NewNamedSemaphoreSM[K comparable](maxCount uint) *NamedSemaphoreSM[K] {
	return &NamedSemaphoreSM[K]{
		maxCount:    maxCount,
		internalMap: xsync.NewMap[K, *Semaphore](),
	}
}

// Release releases one acquisition for the semaphore associated with slug.
//
// Parameters:
//   - slug: key identifying the semaphore to release.
//
// Returns:
//   - error: an error when no matching acquisition exists.
func (nm *NamedSemaphoreSM[K]) Release(slug K) error {
	v2, _ := nm.internalMap.LoadOrStore(slug, NewSemaphore(nm.maxCount))
	return v2.Release()
}

// Acquire waits for and acquires a permit from the semaphore associated with slug.
//
// Parameters:
//   - slug: key identifying the semaphore to acquire.
func (nm *NamedSemaphoreSM[K]) Acquire(slug K) {
	v2, _ := nm.internalMap.LoadOrStore(slug, NewSemaphore(nm.maxCount))
	v2.Acquire()
}
