package observable_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/IlyaFloppy/observable"
)

func TestAPI(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	obj := observable.New[string]("initial")
	require.Equal(t, "initial", obj.Value())

	ch := obj.Subscribe(ctx, true)
	var results []string

	// not using WaitGroup to make sure test passes without "sync" import
	done := make(chan struct{})
	go func() {
		for r := range ch {
			results = append(results, r)
			if len(results) == 4 {
				close(done)
				return
			}
		}
	}()

	obj.Update("value")
	obj.Update("is")
	obj.Update("updated")

	<-done
	cancel()

	require.Equal(t, []string{"initial", "value", "is", "updated"}, results)
}
