package util

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	reloadCalled   = 0
	shutdownCalled = 0
	wg             sync.WaitGroup
)

func test1() {
	reloadCalled++
}

func test2() {
	wg.Done()
	shutdownCalled++
	fmt.Println("====TERMINATED===")
}

func TestTrapAllSignals(t *testing.T) {

	TrapSignals()

	OnReloadEvent(test1)
	OnReloadEvent(test1)
	OnReloadEvent(test1)

	OnShutdownEvent(test2)

	time.Sleep(time.Second)
	wg.Add(1)
	go func() {
		FireReload()
		//FireTerminate()
		time.Sleep(time.Second)
		wg.Done()
	}()

	wg.Wait()

	if reloadCalled != 3 {
		t.Errorf("reload test, expect %d, actual:%d", 3, reloadCalled)
		t.Fail()
	}

	if shutdownCalled != 1 {
		t.Errorf("shutdown test, expect %d, actual:%d", 1, shutdownCalled)
		t.Fail()
	}
}
