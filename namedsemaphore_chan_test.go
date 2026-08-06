package advsync_test

import (
	"testing"
	"time"

	"github.com/vitalick/advsync"
)

type namedSemaphoreChan interface {
	Acquire(string)
	Release(string)
}

func TestNamedSemaphoreChanAcquireBlocksAtCapacity(t *testing.T) {
	tests := map[string]namedSemaphoreChan{
		"map":       advsync.NewNamedSemaphoreChan[string](1),
		"xsync.Map": advsync.NewNamedSemaphoreChanSM[string](1),
	}

	for name, semaphore := range tests {
		t.Run(name, func(t *testing.T) {
			semaphore.Acquire("key")
			acquired := make(chan struct{})
			go func() {
				semaphore.Acquire("key")
				close(acquired)
			}()

			select {
			case <-acquired:
				t.Fatal("acquire succeeded while the named semaphore was at capacity")
			case <-time.After(50 * time.Millisecond):
			}

			semaphore.Release("key")
			select {
			case <-acquired:
			case <-time.After(time.Second):
				t.Fatal("acquire did not continue after release")
			}
			semaphore.Release("key")
		})
	}
}

func TestNamedSemaphoreChanDifferentKeysDoNotShareCapacity(t *testing.T) {
	tests := map[string]namedSemaphoreChan{
		"map":       advsync.NewNamedSemaphoreChan[string](1),
		"xsync.Map": advsync.NewNamedSemaphoreChanSM[string](1),
	}

	for name, semaphore := range tests {
		t.Run(name, func(t *testing.T) {
			semaphore.Acquire("first")
			defer semaphore.Release("first")

			acquired := make(chan struct{})
			go func() {
				semaphore.Acquire("second")
				close(acquired)
			}()

			select {
			case <-acquired:
			case <-time.After(time.Second):
				t.Fatal("a different key remained blocked")
			}
			semaphore.Release("second")
		})
	}
}
