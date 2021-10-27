package main

import (
	"fmt"
	"time"
)

type EventMsg struct {
	Msg string
}

type IObserver interface {
	OnNotify(e EventMsg)
}

type INotifier interface {
	Register(o IObserver)
	Deregister(o IObserver)
	Notify(e EventMsg)
}

type EventObserver struct {
	Id int64
}

func (m *EventObserver) OnNotify(e EventMsg) {
	fmt.Printf("Id:%d OnNotify Msg:%s\n", m.Id, e.Msg)
}

type EventNotifier struct {
	observers map[IObserver]struct{}
}

func (m *EventNotifier) Register(o IObserver) {
	m.observers[o] = struct{}{}
}

func (m *EventNotifier) Deregister(o IObserver) {
	delete(m.observers, o)
}

func (m *EventNotifier) Notify(e EventMsg) {
	for o, _ := range m.observers {
		if o != nil {
			o.OnNotify(e)
		}
	}
}

func main() {
	n := EventNotifier{observers: map[IObserver]struct{}{}}

	o1 := &EventObserver{Id: 1}
	n.Register(o1)
	o2 := &EventObserver{Id: 2}
	n.Register(o2)

	stop := time.NewTimer(5 * time.Second)
	tick1 := time.NewTicker(1 * time.Second)

	for {
		select {
		case date := <-tick1.C:
			n.Notify(EventMsg{Msg: fmt.Sprintf("date:%d", date.Unix())})
		case <-stop.C:
			return
		}
	}

}
