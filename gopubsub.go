package gopubsub

import (
	"sync"
)

const (
	defaultPubSubTopicCapacity = 100000
	initialSubscriberCapacity  = 100
)

// PubSub ...
type PubSub struct {
	mu         *sync.Mutex
	registries []registry
}

type registry struct {
	topic       string
	subscribers []chan interface{}
}

// NewPubSub returns PubSub object
func NewPubSub() *PubSub {
	ps := &PubSub{
		mu:         new(sync.Mutex),
		registries: make([]registry, defaultPubSubTopicCapacity),
	}

	for i := 0; i < defaultPubSubTopicCapacity; i++ {
		ps.registries[i].subscribers = make([]chan interface{}, initialSubscriberCapacity)
	}

	go ps.start()
	return ps
}

func (p *PubSub) start() {
}

// Subscribe ...
func (p *PubSub) Subscribe(topic string) <-chan interface{} {
	ch := make(chan interface{}, 1)

	// TODO generate hash

	return ch
}

// Publish ...
func (p *PubSub) Publish(topic string, message interface{}) {
}
