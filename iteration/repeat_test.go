package iteration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	assert.Equal(t, expected, repeated)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
