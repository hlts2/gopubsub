package gopubsub

import (
	"testing"
)

func TestGenerateHash(t *testing.T) {
	tests := []struct {
		text     string
		expected uint32
	}{
		{
			"abc",
			uint32(1134309195),
		},
	}

	for i, test := range tests {
		got := generateHash(test.text)

		if test.expected != got {
			t.Errorf("tests[%d] - generateHash is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}
