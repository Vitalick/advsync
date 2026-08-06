package advsync_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vitalick/advsync"
)

func TestSemaphoreAcquireBlocksAtCapacity(t *testing.T) {
	semaphore := advsync.NewSemaphore(1)
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

	require.NoError(t, semaphore.Release())
	select {
	case <-acquired:
	case <-time.After(time.Second):
		t.Fatal("acquire did not continue after release")
	}
	require.NoError(t, semaphore.Release())
}

func TestSemaphoreReleaseWithoutAcquireReturnsError(t *testing.T) {
	semaphore := advsync.NewSemaphore(1)

	require.Error(t, semaphore.Release())
}
