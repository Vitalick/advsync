package advsync_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vitalick/advsync"
)

type namedRWMutex interface {
	Lock(string)
	Unlock(string)
	UnlockSafe(string) bool
	RLock(string)
	RUnlock(string)
	RUnlockSafe(string) bool
}

func TestNamedRWMutexSafeUnlocks(t *testing.T) {
	tests := map[string]namedRWMutex{
		"map":       advsync.NewNamedRWMutex[string](),
		"xsync.Map": func() namedRWMutex { mutex := advsync.NewNamedRWMutexSM[string](); return &mutex }(),
	}

	for name, mutex := range tests {
		t.Run(name, func(t *testing.T) {
			require.False(t, mutex.UnlockSafe("key"))
			require.False(t, mutex.RUnlockSafe("key"))

			mutex.Lock("key")
			require.True(t, mutex.UnlockSafe("key"))
			mutex.RLock("key")
			require.True(t, mutex.RUnlockSafe("key"))
		})
	}
}

func TestNamedRWMutexReadersBlockWriter(t *testing.T) {
	tests := map[string]namedRWMutex{
		"map":       advsync.NewNamedRWMutex[string](),
		"xsync.Map": func() namedRWMutex { mutex := advsync.NewNamedRWMutexSM[string](); return &mutex }(),
	}

	for name, mutex := range tests {
		t.Run(name, func(t *testing.T) {
			mutex.RLock("key")
			readerAcquired := make(chan struct{})
			releaseReader := make(chan struct{})
			go func() {
				mutex.RLock("key")
				close(readerAcquired)
				<-releaseReader
				mutex.RUnlock("key")
			}()

			select {
			case <-readerAcquired:
			case <-time.After(time.Second):
				t.Fatal("second reader did not acquire the lock")
			}

			writerAcquired := make(chan struct{})
			go func() {
				mutex.Lock("key")
				close(writerAcquired)
				mutex.Unlock("key")
			}()

			select {
			case <-writerAcquired:
				t.Fatal("writer acquired while readers held the lock")
			case <-time.After(50 * time.Millisecond):
			}

			mutex.RUnlock("key")
			close(releaseReader)

			select {
			case <-writerAcquired:
			case <-time.After(time.Second):
				t.Fatal("writer did not acquire after readers released the lock")
			}
		})
	}
}
