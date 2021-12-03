package advSync

import "sync"

type NamedMutex struct {
	mapLock     sync.RWMutex
	internalMap map[interface{}]*sync.Mutex
}

func (nm *NamedMutex) Unlock(slug interface{}) {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.Mutex{}
		nm.internalMap[slug].Unlock()
		locker = nm.internalMap[slug]
		nm.mapLock.Unlock()
	} else {
		locker.Unlock()
	}
}

func (nm *NamedMutex) Lock(slug interface{}) {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		nm.mapLock.Lock()
		nm.internalMap[slug] = &sync.Mutex{}
		nm.internalMap[slug].Lock()
		locker = nm.internalMap[slug]
		nm.mapLock.Unlock()
	} else {
		locker.Lock()
	}
}
