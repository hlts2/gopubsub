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
