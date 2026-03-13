package advsync_test

import (
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vitalick/advsync"
)

// minimal interface for NamedMutex-like types
// it is satisfied by *advsync.NamedMutex[K] and *advsync.NamedMutexSM[K]
// and represents just Lock/Unlock by key
// Keeping it here avoids changing the library's public API.
type namedLocker[K comparable] interface {
	Lock(K)
	Unlock(K)
}

// shared test logic to verify that locking the same key serializes access
func runSameKeyNoRace[K comparable](t *testing.T, name string, nl namedLocker[K], key K) {
	t.Helper()

	n := 20000
	var wg sync.WaitGroup
	wg.Add(n)

	counter := 0
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			nl.Lock(key)
			// Simulate some work and encourage scheduler switches
			tmp := counter
			runtime.Gosched()
			tmp++
			counter = tmp
			nl.Unlock(key)
		}()
	}
	wg.Wait()

	require.Equal(t, n, counter, "%s: all increments must be observed under the same key lock", name)
}

func TestNamedMutexes_SameKey_NoRace(t *testing.T) {
	// map+RWMutex implementation
	{
		nm := advsync.NewNamedMutex[string]() // returns *NamedMutex[string]
		runSameKeyNoRace[string](t, "NamedMutex", nm, "k")
	}
	// xsync.Map-based implementation
	{
		nmSM := advsync.NewNamedMutexSM[string]()               // returns NamedMutexSM[string] (value)
		runSameKeyNoRace[string](t, "NamedMutexSM", &nmSM, "k") // pass pointer to match interface
	}
}
