package advsync

import (
	"github.com/puzpuzpuz/xsync/v3"
)

// NamedSemaphoreSM is a named semaphore via sync.Map
type NamedSemaphoreSM[K comparable] struct {
	internalMap *xsync.MapOf[K, *Semaphore]
	maxCount    uint
}

// NewNamedSemaphoreSM create new named read/write mutex
func NewNamedSemaphoreSM[K comparable](maxCount uint) *NamedSemaphoreSM[K] {
	return &NamedSemaphoreSM[K]{
		maxCount:    maxCount,
		internalMap: xsync.NewMapOf[K, *Semaphore](),
	}
}

// Release semaphore by name
func (nm *NamedSemaphoreSM[K]) Release(slug K) error {
	v2, _ := nm.internalMap.LoadOrStore(slug, NewSemaphore(nm.maxCount))
	return v2.Release()
}

// Acquire semaphore by name
func (nm *NamedSemaphoreSM[K]) Acquire(slug K) {
	v2, _ := nm.internalMap.LoadOrStore(slug, NewSemaphore(nm.maxCount))
	v2.Acquire()
}
