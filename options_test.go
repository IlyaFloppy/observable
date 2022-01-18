package observable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithSendCurrent(t *testing.T) {
	o := params{}

	require.False(t, o.withCurrent)

	o = WithCurrent(true)(o)
	require.True(t, o.withCurrent)
}

func TestWithBufferSize(t *testing.T) {
	o := params{}

	require.Equal(t, 0, o.bufferSize)

	o = WithBuffer(42)(o)
	require.Equal(t, 42, o.bufferSize)
}
