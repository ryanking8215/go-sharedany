package sharedany

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SharedBytes(t *testing.T) {
	create := func() []byte {
		return make([]byte, 1000)
	}
	done := false
	sbs := New[[]byte](create(), 2, func(s *Shared[[]byte]) {
		done = true
	})
	assert.Len(t, sbs.Data(), 1000)
	sbs.Done()
	assert.False(t, done)
	sbs.Done()
	assert.True(t, done)
}
