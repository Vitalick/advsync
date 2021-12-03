package advSync

import "sync"

type NamedMutexSM struct {
	internalMap sync.Map
}

func (nm *NamedMutexSM) Unlock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.Mutex{})
	v2.(*sync.Mutex).Unlock()
}

func (nm *NamedMutexSM) Lock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.Mutex{})
	v2.(*sync.Mutex).Lock()
}
