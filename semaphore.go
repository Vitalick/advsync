package advsync

import (
	"errors"
	"sync"
)

//Semaphore is a semaphore primitive based on sync.Cond
type Semaphore struct {
	cond     *sync.Cond
	counter  uint
	maxCount uint
}

//NewSemaphore return new Semaphore with max count of acquiries
func NewSemaphore(maxCount uint) *Semaphore {
	cond := sync.NewCond(&sync.Mutex{})
	return &Semaphore{
		cond:     cond,
		maxCount: maxCount,
	}
}

//Acquire waiting until counter bigger than max count
func (s *Semaphore) Acquire() {
	s.cond.L.Lock()
	if s.counter >= s.maxCount {
		s.cond.Wait()
	}
	s.counter++
	s.cond.L.Unlock()
}

//Release decrease counter until counter bigger than 0
func (s *Semaphore) Release() error {
	s.cond.L.Lock()
	if s.counter < 1 {
		s.cond.L.Unlock()
		return errors.New("not found acquire")
	}
	s.counter--
	s.cond.L.Unlock()
	s.cond.Broadcast()
	return nil
}
