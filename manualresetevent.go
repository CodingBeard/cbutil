package cbutil

import "sync"

type ManualResetEvent struct {
	lock   *sync.Mutex
	locked bool
	wait   chan int
}

func NewManualResetEvent(locked bool) *ManualResetEvent {
	event := ManualResetEvent{
		lock:   &sync.Mutex{},
		locked: locked,
		wait:   make(chan int),
	}

	return &event
}

func (e *ManualResetEvent) Wait() {
	if e.locked {
		<-e.wait
	}
}

func (e *ManualResetEvent) Set() {
	e.lock.Lock()
	e.locked = false
	e.lock.Unlock()

	e.wait <- 1
}

func (e *ManualResetEvent) Reset() {
	e.lock.Lock()
	e.locked = true
	e.lock.Unlock()
}
