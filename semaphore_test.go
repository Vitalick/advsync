package advsync_test

import (
	"fmt"
	advsync "github.com/vitalick/adv-sync"
	"testing"
	"time"
)

var sema = advsync.NewSemaphore(5)

func PrintThread(c int) {
	sema.Acquire()
	for range make([]struct{}, 5) {
		fmt.Println(c)
		time.Sleep(time.Millisecond * 500)
	}
	if err := sema.Release(); err != nil {
		panic(err)
	}
}

func TestSemaphore_Semaphore(t *testing.T) {
	for i := range make([]struct{}, 10) {
		go PrintThread(i)
	}
	time.Sleep(time.Second * 10)
}
