package advsync

import "sync"

//NamedSemaphoreSM is a named semaphore via sync.Map
type NamedSemaphoreSM struct {
	internalMap sync.Map
	maxCount    uint
}

//NewNamedSemaphoreSM create new named read/write mutex
func NewNamedSemaphoreSM(maxCount uint) *NamedSemaphoreSM {
	return &NamedSemaphoreSM{
		maxCount: maxCount,
	}
}

//Release semaphore by name
func (nm *NamedSemaphoreSM) Release(slug interface{}) error {
	v2, _ := nm.internalMap.LoadOrStore(slug, NewSemaphore(nm.maxCount))
	return v2.(*Semaphore).Release()
}

//Acquire semaphore by name
func (nm *NamedSemaphoreSM) Acquire(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, NewSemaphore(nm.maxCount))
	v2.(*Semaphore).Acquire()
}
