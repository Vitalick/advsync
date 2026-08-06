package advsync

// SemaphoreChan is a counting semaphore implemented with a channel.
type SemaphoreChan struct {
	ch chan struct{}
}

// NewSemaphoreChan creates a SemaphoreChan with at most maxCount concurrent acquisitions.
//
// Parameters:
//   - maxCount: maximum number of acquired permits.
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

// Acquire waits until a permit is available and then acquires it.
func (s *SemaphoreChan) Acquire() {
	s.ch <- struct{}{}
}

// Release releases one acquired permit.
func (s *SemaphoreChan) Release() {
	defer func() { <-s.ch }()
}

// Close closes the underlying channel.
//
// Cases:
//   - Do not call Close while another goroutine can acquire or release a permit.
func (s *SemaphoreChan) Close() {
	close(s.ch)
}
