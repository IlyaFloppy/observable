package observable

import (
	"context"
	"sync"
)

type Object[T any] struct {
	lock  sync.Mutex
	state *state[T]
}

func (o *Object[T]) Get() T {
	o.lock.Lock()
	defer o.lock.Unlock()

	return o.state.value
}

func (o *Object[T]) Set(value T) {
	o.lock.Lock()
	defer o.lock.Unlock()

	o.state = o.state.update(value)
}

func (o *Object[T]) Subscribe(ctx context.Context, sendCurrent bool) <-chan T {
	o.lock.Lock()
	defer o.lock.Unlock()

	s := o.state
	ch := make(chan T)

	go func() {
		defer close(ch)

		if sendCurrent {
			select {
			case <-ctx.Done():
				return
			case ch <- s.value:
			}
		}

		for {
			select {
			case <-ctx.Done():
				return
			case <-s.done:
				s = s.next
			}

			select {
			case <-ctx.Done():
				return
			case ch <- s.value:
			}
		}
	}()

	return ch
}
