package sharedany

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BytesPool(t *testing.T) {
	create := func() []byte {
		return make([]byte, 1000)
	}

	put := false
	bspool := NewPool[[]byte](create, func(sbs *Shared[[]byte]) {
		put = true
	})
	sbs := bspool.Get()
	assert.Len(t, sbs.Data(), 1000)
	sbs.Add(3)
	sbs.Done()
	assert.False(t, put)
	sbs.Done()
	assert.False(t, put)
	sbs.Done()
	assert.True(t, put)
}
