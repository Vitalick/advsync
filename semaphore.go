package advsync

import (
	"errors"
	"sync"
)

// Semaphore is a counting semaphore implemented with sync.Cond.
type Semaphore struct {
	cond     *sync.Cond
	counter  uint
	maxCount uint
}

// NewSemaphore creates a Semaphore with at most maxCount concurrent acquisitions.
//
// Parameters:
//   - maxCount: maximum number of acquired permits.
func NewSemaphore(maxCount uint) *Semaphore {
	cond := sync.NewCond(&sync.Mutex{})
	return &Semaphore{
		cond:     cond,
		maxCount: maxCount,
	}
}

// Acquire waits until a permit is available and then acquires it.
func (s *Semaphore) Acquire() {
	s.cond.L.Lock()
	if s.counter >= s.maxCount {
		s.cond.Wait()
	}
	s.counter++
	s.cond.L.Unlock()
}

// Release releases one acquired permit.
//
// Returns:
//   - error: an error when there is no acquisition to release.
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
