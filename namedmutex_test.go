package advsync_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vitalick/advsync"
)

type namedMutex interface {
	Lock(string)
	Unlock(string)
	UnlockSafe(string) bool
}

func TestNamedMutexUnlockSafe(t *testing.T) {
	t.Parallel()

	tests := map[string]namedMutex{
		"map":       advsync.NewNamedMutex[string](),
		"xsync.Map": func() namedMutex { mutex := advsync.NewNamedMutexSM[string](); return &mutex }(),
	}

	for name, mutex := range tests {
		t.Run(name, func(t *testing.T) {
			require.False(t, mutex.UnlockSafe("key"))
			mutex.Lock("key")
			require.True(t, mutex.UnlockSafe("key"))
		})
	}
}

func TestNamedMutexDifferentKeysDoNotBlock(t *testing.T) {
	tests := map[string]namedMutex{
		"map":       advsync.NewNamedMutex[string](),
		"xsync.Map": func() namedMutex { mutex := advsync.NewNamedMutexSM[string](); return &mutex }(),
	}

	for name, mutex := range tests {
		t.Run(name, func(t *testing.T) {
			mutex.Lock("first")
			defer mutex.Unlock("first")

			acquired := make(chan struct{})
			go func() {
				mutex.Lock("second")
				defer mutex.Unlock("second")
				close(acquired)
			}()

			select {
			case <-acquired:
			case <-time.After(time.Second):
				t.Fatal("a different key remained blocked")
			}
		})
	}
}

func TestNamedMutexSameKeyBlocks(t *testing.T) {
	tests := map[string]namedMutex{
		"map":       advsync.NewNamedMutex[string](),
		"xsync.Map": func() namedMutex { mutex := advsync.NewNamedMutexSM[string](); return &mutex }(),
	}

	for name, mutex := range tests {
		t.Run(name, func(t *testing.T) {
			mutex.Lock("key")
			acquired := make(chan struct{})
			go func() {
				mutex.Lock("key")
				close(acquired)
				mutex.Unlock("key")
			}()

			select {
			case <-acquired:
				t.Fatal("the same key acquired a locked mutex")
			case <-time.After(50 * time.Millisecond):
			}

			mutex.Unlock("key")
			select {
			case <-acquired:
			case <-time.After(time.Second):
				t.Fatal("the same key did not acquire after unlock")
			}
		})
	}
}
