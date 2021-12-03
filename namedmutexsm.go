package advSync

import "sync"

//NamedMutexSM is a named mutex via sync.Map
type NamedMutexSM struct {
	internalMap sync.Map
}

//Unlock mutex by name
func (nm *NamedMutexSM) Unlock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.Mutex{})
	v2.(*sync.Mutex).Unlock()
}

//Lock mutex by name
func (nm *NamedMutexSM) Lock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.Mutex{})
	v2.(*sync.Mutex).Lock()
}
