package advSync

import "sync"

//NamedRWMutexSM is a named read/write mutex via sync.Map
type NamedRWMutexSM struct {
	internalMap sync.Map
}

//Unlock mutex by name
func (nm *NamedRWMutexSM) Unlock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	v2.(*sync.RWMutex).Unlock()
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

//RLock mutex by name
func (nm *NamedRWMutexSM) RLock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	v2.(*sync.RWMutex).RLock()
}
