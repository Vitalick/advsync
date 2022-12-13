package advsync

// SemaphoreChan is a semaphore primitive based on chan
type SemaphoreChan struct {
	ch chan struct{}
}

// NewSemaphoreChan return new Semaphore with max count of acquiries
func NewSemaphoreChan(maxCount uint) *SemaphoreChan {
	if maxCount == 0 {
		return &SemaphoreChan{
			ch: make(chan struct{}),
		}
	}
	return &SemaphoreChan{
		ch: make(chan struct{}, maxCount),
	}
}

// Acquire waiting until counter bigger than max count
func (s *SemaphoreChan) Acquire() {
	s.ch <- struct{}{}
}

// Release decrease counter until counter bigger than 0
func (s *SemaphoreChan) Release() {
	defer func() { <-s.ch }()
}

func (s *SemaphoreChan) Close() {
	close(s.ch)
}
