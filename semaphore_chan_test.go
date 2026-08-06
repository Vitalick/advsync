package advsync_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/vitalick/advsync"
)

func TestSemaphoreChanAcquireBlocksAtCapacity(t *testing.T) {
	semaphore := advsync.NewSemaphoreChan(1)
	semaphore.Acquire()

	acquired := make(chan struct{})
	go func() {
		semaphore.Acquire()
		close(acquired)
	}()

	select {
	case <-acquired:
		t.Fatal("acquire succeeded while the semaphore was at capacity")
	case <-time.After(50 * time.Millisecond):
	}

	semaphore.Release()
	select {
	case <-acquired:
	case <-time.After(time.Second):
		t.Fatal("acquire did not continue after release")
	}
	semaphore.Release()
}

func TestSemaphoreChanZeroCapacityHandoff(t *testing.T) {
	semaphore := advsync.NewSemaphoreChan(0)
	acquired := make(chan struct{})

	go func() {
		semaphore.Acquire()
		close(acquired)
	}()

	semaphore.Release()
	select {
	case <-acquired:
	case <-time.After(time.Second):
		t.Fatal("zero-capacity semaphore did not hand off a permit")
	}
}

func TestSemaphoreChanClosePreventsFurtherAcquire(t *testing.T) {
	semaphore := advsync.NewSemaphoreChan(1)
	semaphore.Close()

	require.Panics(t, func() {
		semaphore.Acquire()
	})
}
