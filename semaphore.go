package advSync

import (
	"errors"
	"sync"
)

type Semaphore struct {
	cond     *sync.Cond
	counter  int
	maxCount int
}

func NewSemaphore(maxCount int) *Semaphore {
	cond := sync.NewCond(&sync.Mutex{})
	return &Semaphore{
		cond:     cond,
		maxCount: maxCount,
	}
}

func (s *Semaphore) Acquire() {
	s.cond.L.Lock()
	if s.counter >= s.maxCount {
		s.cond.Wait()
	}
	s.counter++
	s.cond.L.Unlock()
}

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
