package advsync

import (
	"reflect"
	"sync"
)

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
		nm.mapLock.Unlock()
		return
	}
	locker.Unlock()
}

//UnlockSafe mutex by name
func (nm *NamedRWMutex) UnlockSafe(slug interface{}) bool {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		return false
	}
	state := reflect.ValueOf(locker).Elem().FieldByName("w").FieldByName("state")
	vb := state.Int()&mutexLocked == mutexLocked
	if !vb {
		return false
	}
	locker.Unlock()
	return true
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
		nm.mapLock.Unlock()
		return
	}
	locker.Lock()
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
		nm.mapLock.Unlock()
		return
	}
	locker.RUnlock()

}

//RUnlockSafe mutex by name
func (nm *NamedRWMutex) RUnlockSafe(slug interface{}) bool {
	nm.mapLock.RLock()
	locker, ok := nm.internalMap[slug]
	nm.mapLock.RUnlock()
	if !ok {
		return false
	}
	state := reflect.ValueOf(locker).Elem().FieldByName("readerCount")
	vb := state.Int() > 0
	if !vb {
		return false
	}
	locker.RUnlock()
	return true
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
		nm.mapLock.Unlock()
		return
	}
	locker.RLock()

}
