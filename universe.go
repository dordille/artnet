package artnet

import (
	"sync"
	"time"
)

type Universe struct {
	node *Node

	universe uint8
	channels [512]uint8

	dirty bool
	mutex sync.Mutex
}

func NewUniverse(i uint8, n *Node, r time.Duration) *Universe {
	u := &Universe{
		node:     n,
		universe: i,
	}

	go func() {
		for {
			time.Sleep(r)
			u.Send()
		}
	}()

	return u
}

func (u *Universe) Set(ch int, v uint8) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	u.channels[ch] = v
	u.dirty = true
}

func (u *Universe) MultiSet(ch int, v []uint8) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	copy(u.channels[ch:], v)
	u.dirty = true
}

func (u *Universe) ClearMultiSet(ch int, v []uint8) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	for k, _ := range u.channels {
		u.channels[k] = 0
	}

	copy(u.channels[ch:], v)
	u.dirty = true
}

func (u *Universe) Send() {
	if u.dirty == false {
		return
	}

	u.mutex.Lock()
	defer u.mutex.Unlock()

	u.node.Dmx(u.universe, u.channels)
	u.dirty = false
}
