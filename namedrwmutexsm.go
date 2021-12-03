package advSync

import "sync"

type NamedRWMutexSM struct {
	internalMap sync.Map
}

func (nm *NamedRWMutexSM) Unlock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	v2.(*sync.RWMutex).Unlock()
}

func (nm *NamedRWMutexSM) Lock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	v2.(*sync.RWMutex).Lock()
}

func (nm *NamedRWMutexSM) RUnlock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	v2.(*sync.RWMutex).RUnlock()
}

func (nm *NamedRWMutexSM) RLock(slug interface{}) {
	v2, _ := nm.internalMap.LoadOrStore(slug, &sync.RWMutex{})
	v2.(*sync.RWMutex).RLock()
}
