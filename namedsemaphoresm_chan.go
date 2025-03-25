package advsync

import (
	"github.com/puzpuzpuz/xsync/v3"
)

// NamedSemaphoreChanSM is a named semaphore via sync.Map
type NamedSemaphoreChanSM[K comparable] struct {
	internalMap *xsync.MapOf[K, *SemaphoreChan]
	maxCount    uint
}

// NewNamedSemaphoreChanSM create new named read/write mutex
func NewNamedSemaphoreChanSM[K comparable](maxCount uint) *NamedSemaphoreChanSM[K] {
	return &NamedSemaphoreChanSM[K]{
		maxCount:    maxCount,
		internalMap: xsync.NewMapOf[K, *SemaphoreChan](),
	}
}

// Release semaphore by name
func (nm *NamedSemaphoreChanSM[K]) Release(slug K) {
	v2, _ := nm.internalMap.LoadOrStore(slug, NewSemaphoreChan(nm.maxCount))
	v2.Release()
}

// Acquire semaphore by name
func (nm *NamedSemaphoreChanSM[K]) Acquire(slug K) {
	v2, _ := nm.internalMap.LoadOrStore(slug, NewSemaphoreChan(nm.maxCount))
	v2.Acquire()
}
