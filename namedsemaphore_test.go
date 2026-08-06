package advsync_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vitalick/advsync"
)

type namedSemaphore interface {
	Acquire(string)
	Release(string) error
}

func TestNamedSemaphoreAcquireBlocksAtCapacity(t *testing.T) {
	tests := map[string]namedSemaphore{
		"map":       advsync.NewNamedSemaphore[string](1),
		"xsync.Map": advsync.NewNamedSemaphoreSM[string](1),
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

			require.NoError(t, semaphore.Release("key"))
			select {
			case <-acquired:
			case <-time.After(time.Second):
				t.Fatal("acquire did not continue after release")
			}
			require.NoError(t, semaphore.Release("key"))
		})
	}
}

func TestNamedSemaphoreDifferentKeysDoNotShareCapacity(t *testing.T) {
	tests := map[string]namedSemaphore{
		"map":       advsync.NewNamedSemaphore[string](1),
		"xsync.Map": advsync.NewNamedSemaphoreSM[string](1),
	}

	for name, semaphore := range tests {
		t.Run(name, func(t *testing.T) {
			semaphore.Acquire("first")
			defer func() { require.NoError(t, semaphore.Release("first")) }()

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
			require.NoError(t, semaphore.Release("second"))
		})
	}
}
