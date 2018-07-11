package gopubsub

import (
	"testing"
)

func TestNew(t *testing.T) {
	ps := NewPubSub()
	if ps == nil {
		t.Errorf("NewPubSub is nil")
	}

	got := len(ps.registries)
	if defaultPubSubTopicCapacity != got {
		t.Errorf("topic capacity is wrong. expected: %v, got: %v", defaultPubSubTopicCapacity, got)
	}

	for _, registory := range ps.registries {
		if "" != registory.topic {
			t.Errorf("topic word is wrong. expected: %v, got: %v", "", registory.topic)
		}
	}
}

func TestSubscribe(t *testing.T) {
	ps := NewPubSub()

	subscriber1 := ps.Subscribe("t1")
	subscriber2 := ps.Subscribe("t1")
	subscriber3 := ps.Subscribe("t2")

	ps.Publish("t1", "hi")
	ps.Publish("t2", "hello")

	got := <-subscriber1.Read()
	if "hi" != got {
		t.Errorf("message of subscriber1 is wrong. expected: %v, got: %v", "hi", got)
	}

	got = <-subscriber2.Read()
	if "hi" != got {
		t.Errorf("message of subscriber2 is wrong. expected: %v, got: %v", "hi", got)
	}

	got = <-subscriber3.Read()
	if "hello" != got {
		t.Errorf("message of subscribers is wrong. expected: %v, got: %v", "hello", got)
	}
}
