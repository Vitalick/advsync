package advsync

import "sync"

// NamedSemaphoreChanSM is a named semaphore via sync.Map
type NamedSemaphoreChanSM struct {
	internalMap sync.Map
	maxCount    uint
}

// NewNamedSemaphoreChanSM create new named read/write mutex
func NewNamedSemaphoreChanSM(maxCount uint) *NamedSemaphoreChanSM {
	return &NamedSemaphoreChanSM{
		maxCount: maxCount,
	}
}

// Release semaphore by name
func (nm *NamedSemaphoreChanSM) Release(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, NewSemaphoreChan(nm.maxCount))
	v2.(*SemaphoreChan).Release()
}

// Acquire semaphore by name
func (nm *NamedSemaphoreChanSM) Acquire(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, NewSemaphoreChan(nm.maxCount))
	v2.(*SemaphoreChan).Acquire()
}
