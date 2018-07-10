package gopubsub

import (
	"testing"
)

func TestNewPubSub(t *testing.T) {
	ps := NewPubSub()

	if ps == nil {
		t.Errorf("NewPubSub is nil")
	}

	got := len(ps.registries)
	if defaultPubSubTopicCapacity != got {
		t.Errorf("PubSub capacity is wrong. expected: %v, got: %v", defaultPubSubTopicCapacity, got)
	}
}
