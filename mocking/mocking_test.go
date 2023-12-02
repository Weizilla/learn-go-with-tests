package mocking

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMocking(t *testing.T) {
	t.Run("mocking", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}

		Countdown(buffer, sleeper)

		actual := buffer.String()
		expected :=
			`3
2
1
Go!`

		assert.Equal(t, expected, actual)
		assert.Equal(t, 3, sleeper.Calls)
	})
}
