package observable_test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/IlyaFloppy/observable"
)

func TestSubscribe(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	obj := observable.New("initial")
	require.Equal(t, "initial", obj.Get())

	ch := obj.Subscribe(ctx, observable.WithCurrent(true))

	var results []string

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		for r := range ch {
			results = append(results, r)
			wg.Done()
		}
	}()

	obj.Set("value")
	obj.Set("is")
	obj.Set("updated")

	wg.Wait()
	cancel()

	require.Equal(t, []string{"initial", "value", "is", "updated"}, results)
}

func TestSubscribeSendCurrentContext(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			obj := observable.New("initial")
			ch := obj.Subscribe(ctx, observable.WithCurrent(true))

			if val, ok := <-ch; ok {
				// select has chosen send branch
				require.Equal(t, "initial", val)
			} else {
				// select has chosen ctx.Done() branch
			}
		})
	}
}

func TestSubscribeContext(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			obj := observable.New("initial")
			ch := obj.Subscribe(ctx, observable.WithCurrent(false))
			obj.Set("value")

			if val, ok := <-ch; ok {
				// passed through both selects
				require.Equal(t, "value", val)
			} else {
				// one of selects has chosen ctx.Done() branch
			}
		})
	}
}
