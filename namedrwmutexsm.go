package advsync

import (
	"reflect"
	"sync"
)

//NamedRWMutexSM is a named read/write mutex via sync.Map
type NamedRWMutexSM struct {
	internalMap sync.Map
}

//Unlock mutex by name
func (nm *NamedRWMutexSM) Unlock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	v2.(*sync.RWMutex).Unlock()
}

//UnlockSafe mutex by name
func (nm *NamedRWMutexSM) UnlockSafe(slug interface{}) bool {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	state := reflect.ValueOf(v2).Elem().FieldByName("w").FieldByName("state")
	vb := state.Int()&mutexLocked == mutexLocked
	if !vb {
		return false
	}
	v2.(*sync.RWMutex).Unlock()
	return true
}

//Lock mutex by name
func (nm *NamedRWMutexSM) Lock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	v2.(*sync.RWMutex).Lock()
}

//RUnlock mutex by name
func (nm *NamedRWMutexSM) RUnlock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	v2.(*sync.RWMutex).RUnlock()
}

//RUnlockSafe mutex by name
func (nm *NamedRWMutexSM) RUnlockSafe(slug interface{}) bool {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	state := reflect.ValueOf(v2).Elem().FieldByName("readerCount")
	vb := state.Int() > 0
	if !vb {
		return false
	}
	v2.(*sync.RWMutex).Unlock()
	return true
}

//RLock mutex by name
func (nm *NamedRWMutexSM) RLock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	v2.(*sync.RWMutex).RLock()
}
