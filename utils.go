package advsync

import (
	"reflect"
	"sync"
)

func unlockSafe(mutex *sync.Mutex) bool {
	reflectValue := reflect.ValueOf(mutex).Elem()
	state := reflectValue.FieldByName("state")
	if !state.IsValid() {
		state = reflectValue.FieldByName("mu").FieldByName("state")
		if !state.IsValid() {
			return false
		}
	}
	vb := state.Int()&mutexLocked == mutexLocked
	if !vb {
		return false
	}
	mutex.Unlock()
	return true
}

func unlockSafeRW(mutexRW *sync.RWMutex) bool {
	reflectValue := reflect.ValueOf(mutexRW).Elem().FieldByName("w")
	state := reflectValue.FieldByName("state")
	if !state.IsValid() {
		state = reflectValue.FieldByName("mu").FieldByName("state")
		if !state.IsValid() {
			return false
		}
	}
	vb := state.Int()&mutexLocked == mutexLocked
	if !vb {
		return false
	}
	mutexRW.Unlock()
	return true
}

func rUnlockSafeRW(mutexRW *sync.RWMutex) bool {
	state := reflect.ValueOf(mutexRW).Elem().FieldByName("readerCount")
	vb := state.Int() > 0
	if !vb {
		return false
	}
	mutexRW.RUnlock()
	return true
}
