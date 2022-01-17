package observable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithSendCurrent(t *testing.T) {
	o := params{}

	require.False(t, o.sendCurrent)

	o = WithSendCurrent(true)(o)
	require.True(t, o.sendCurrent)
}

func TestWithBufferSize(t *testing.T) {
	o := params{}

	require.Equal(t, 0, o.bufferSize)

	o = WithBufferSize(42)(o)
	require.Equal(t, 42, o.bufferSize)
}
