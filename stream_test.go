package observable_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/IlyaFloppy/observable"
)

func TestStreamChanges(t *testing.T) {
	obj := observable.New("")
	s := obj.Stream()
	obj.Set("initial")

	vals := []string{"first", "second", "third"}
	got := make([]string, 0, 4)

loop:
	for {
		select {
		case <-s.Changes():
			require.True(t, s.HasNext())
			val := s.Next()
			got = append(got, val)

			if len(vals) > 0 {
				obj.Set(vals[0])
				vals = vals[1:]
			} else {
				break loop
			}
		}
	}

	require.Equal(t, []string{"initial", "first", "second", "third"}, got)
}

func TestStreamWaitNext(t *testing.T) {
	obj := observable.New("")
	s := obj.Stream()
	obj.Set("initial")
	obj.Set("first")
	obj.Set("second")
	obj.Set("third")

	got := make([]string, 0, 4)

	for i := 0; i < 4; i++ {
		s.WaitNext()
		val := s.Next()
		got = append(got, val)
	}

	require.Equal(t, []string{"initial", "first", "second", "third"}, got)
}

func TestStreamHasNext(t *testing.T) {
	obj := observable.New("")
	s := obj.Stream()

	require.False(t, s.HasNext())

	obj.Set("initial")
	require.True(t, s.HasNext())

	require.Equal(t, "initial", s.Next())
}
