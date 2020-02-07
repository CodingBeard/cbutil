package cbutil

import (
	"time"
)

type RepeatingTask struct {
	Sleep              time.Duration
	SleepFirst         bool
	ExecuteInGoroutine bool
	Blocking           bool
	Run                func()
}

var RepeatingTaskRecoverFunc func()

func (t RepeatingTask) Start() {
	if RepeatingTaskRecoverFunc == nil {
		RepeatingTaskRecoverFunc = func() {}
	}
	runTask := func() {
		for true {
			if t.SleepFirst {
				Sleep(t.Sleep)
			}
			if t.ExecuteInGoroutine {
				go func() {
					if RepeatingTaskRecoverFunc != nil {
						defer RepeatingTaskRecoverFunc()
					}
					t.Run()
				}()
			} else {
				t.Run()
			}
			if !t.SleepFirst {
				Sleep(t.Sleep)
			}
		}
	}

	if t.Blocking {
		runTask()
	} else {
		go func() {
			if RepeatingTaskRecoverFunc != nil {
				defer RepeatingTaskRecoverFunc()
			}
			runTask()
		}()
	}
}
