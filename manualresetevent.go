package cbutil

import "sync"

type ManualResetEvent struct {
	lock        *sync.Mutex
	locked      bool
	wait        chan int
	waitersLock *sync.Mutex
	waiters     int
}

func NewManualResetEvent(locked bool) *ManualResetEvent {
	event := ManualResetEvent{
		lock:        &sync.Mutex{},
		waitersLock: &sync.Mutex{},
		locked:      locked,
		wait:        make(chan int),
	}

	return &event
}

func (e *ManualResetEvent) Wait() {
	if e.locked {
		e.waitersLock.Lock()
		e.waiters++
		e.waitersLock.Unlock()
		<-e.wait
	}
}

func (e *ManualResetEvent) Set() {
	e.lock.Lock()
	e.locked = false
	e.lock.Unlock()

	for i := 0; i < e.waiters; i++ {
		e.wait <- 1
	}
}

func (e *ManualResetEvent) Reset() {
	e.lock.Lock()
	e.locked = true
	e.lock.Unlock()
}

type MultiManualResetEvent struct {
	events    map[string]*ManualResetEvent
	eventLock *sync.Mutex
}

func NewMultiManualResetEvent() *MultiManualResetEvent {
	return &MultiManualResetEvent{
		eventLock: &sync.Mutex{},
		events:    make(map[string]*ManualResetEvent),
	}
}

func (m *MultiManualResetEvent) Get(key string, locked bool) (isNew bool, event *ManualResetEvent) {
	m.eventLock.Lock()
	defer m.eventLock.Unlock()
	if event, ok := m.events[key]; ok {
		return false, event
	}

	m.events[key] = NewManualResetEvent(locked)

	return true, m.events[key]
}

func (m *MultiManualResetEvent) Remove(key string) {
	m.eventLock.Lock()
	delete(m.events, key)
	m.eventLock.Unlock()
}
