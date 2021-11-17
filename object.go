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

func (o *Object[T]) Subscribe(ctx context.Context, options ...Option) <-chan T {
	o.lock.Lock()
	defer o.lock.Unlock()

	p := params{}
	for _, opt := range options {
		p = opt(p)
	}

	s := o.state
	ch := make(chan T, p.bufferSize)

	go func() {
		defer close(ch)

		if p.sendCurrent {
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

func (o *Object[T]) Stream() Stream[T] {
	return Stream[T]{
		state: o.state,
	}
}
