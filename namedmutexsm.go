package advsync

import (
	"reflect"
	"sync"
)

//NamedMutexSM is a named mutex via sync.Map
type NamedMutexSM struct {
	internalMap sync.Map
}

//Unlock mutex by name
func (nm *NamedMutexSM) Unlock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.Mutex{})
	v2.(*sync.Mutex).Unlock()
}

//UnlockSafe mutex by name
func (nm *NamedMutexSM) UnlockSafe(slug interface{}) bool {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	state := reflect.ValueOf(v2).Elem().FieldByName("state")
	vb := state.Int()&mutexLocked == mutexLocked
	if !vb {
		return false
	}
	v2.(*sync.RWMutex).Unlock()
	return true
}

//Lock mutex by name
func (nm *NamedMutexSM) Lock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.Mutex{})
	v2.(*sync.Mutex).Lock()
}
