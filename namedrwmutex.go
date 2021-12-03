package advSync

import "sync"

type NamedRWMutex struct {
	mapLock     sync.RWMutex
	internalMap map[interface{}]*sync.RWMutex
}

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
