package advsync

import (
	"github.com/puzpuzpuz/xsync/v4"
)

// NamedSemaphoreChanSM provides a channel-based semaphore for each key using xsync.Map.
type NamedSemaphoreChanSM[K comparable] struct {
	internalMap *xsync.Map[K, *SemaphoreChan]
	maxCount    uint
}

// NewNamedSemaphoreChanSM creates a keyed channel-based semaphore collection.
//
// Parameters:
//   - maxCount: maximum simultaneous acquisitions for each key.
func NewNamedSemaphoreChanSM[K comparable](maxCount uint) *NamedSemaphoreChanSM[K] {
	return &NamedSemaphoreChanSM[K]{
		maxCount:    maxCount,
		internalMap: xsync.NewMap[K, *SemaphoreChan](),
	}
}

// Release releases one acquisition for the channel-based semaphore associated with slug.
//
// Parameters:
//   - slug: key identifying the semaphore to release.
func (nm *NamedSemaphoreChanSM[K]) Release(slug K) {
	v2, _ := nm.internalMap.LoadOrStore(slug, NewSemaphoreChan(nm.maxCount))
	v2.Release()
}

// Acquire waits for and acquires a permit from the semaphore associated with slug.
//
// Parameters:
//   - slug: key identifying the semaphore to acquire.
func (nm *NamedSemaphoreChanSM[K]) Acquire(slug K) {
	v2, _ := nm.internalMap.LoadOrStore(slug, NewSemaphoreChan(nm.maxCount))
	v2.Acquire()
}
