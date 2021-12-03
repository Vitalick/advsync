package adv_sync

import "sync"

//NamedRWMutex is a named read/write mutex via sync.RWMutex
type NamedRWMutex struct {
	mapLock     sync.RWMutex
	internalMap map[interface{}]*sync.RWMutex
}

//NewNamedRWMutex create new named read/write mutex
func NewNamedRWMutex() *NamedRWMutex {
	return &NamedRWMutex{
		internalMap: map[interface{}]*sync.RWMutex{},
	}
}

//Unlock mutex by name
func (nm *NamedRWMutex) Unlock(slug interface{}) {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.RWMutex{}
		nm.internalMap[slug].Unlock()
		locker = nm.internalMap[slug]
		nm.mapLock.Unlock()
	} else {
		locker.Unlock()
	}
}

//Lock mutex by name
func (nm *NamedRWMutex) Lock(slug interface{}) {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.RWMutex{}
		nm.internalMap[slug].Lock()
		locker = nm.internalMap[slug]
		nm.mapLock.Unlock()
	} else {
		locker.Lock()
	}
}

//RUnlock mutex by name
func (nm *NamedRWMutex) RUnlock(slug interface{}) {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.RWMutex{}
		nm.internalMap[slug].RUnlock()
		locker = nm.internalMap[slug]
		nm.mapLock.Unlock()
	} else {
		locker.Unlock()
	}
}

//RLock mutex by name
func (nm *NamedRWMutex) RLock(slug interface{}) {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.RWMutex{}
		nm.internalMap[slug].RLock()
		locker = nm.internalMap[slug]
		nm.mapLock.Unlock()
	} else {
		locker.Lock()
	}
}
