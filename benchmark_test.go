package advsync_test

import (
	"testing"

	"github.com/vitalick/advsync"
)

func BenchmarkNamedMutex(b *testing.B) {
	mutex := advsync.NewNamedMutex[string]()
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		mutex.Lock("key")
		mutex.Unlock("key")
	}
}

func BenchmarkNamedMutexSM(b *testing.B) {
	mutex := advsync.NewNamedMutexSM[string]()
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		mutex.Lock("key")
		mutex.Unlock("key")
	}
}

func BenchmarkNamedRWMutex(b *testing.B) {
	mutex := advsync.NewNamedRWMutex[string]()
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		mutex.Lock("key")
		mutex.Unlock("key")
	}
}

func BenchmarkNamedRWMutexSM(b *testing.B) {
	mutex := advsync.NewNamedRWMutexSM[string]()
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		mutex.Lock("key")
		mutex.Unlock("key")
	}
}

func BenchmarkSemaphore(b *testing.B) {
	semaphore := advsync.NewSemaphore(1)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		semaphore.Acquire()
		if err := semaphore.Release(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSemaphoreChan(b *testing.B) {
	semaphore := advsync.NewSemaphoreChan(1)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		semaphore.Acquire()
		semaphore.Release()
	}
}
