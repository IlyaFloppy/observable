package observable

import (
	"context"
	"sync"
)

// New creates a new observable object.
func New[T any](value T) *Object[T] {
	return &Object[T]{state: newState(value)}
}

// Object is a wrapper for any value that allows subscribing for it's changes.
type Object[T any] struct {
	lock  sync.Mutex
	state *state[T]
}

// Get returns current value of an object.
func (o *Object[T]) Get() T {
	o.lock.Lock()
	defer o.lock.Unlock()

	return o.state.value
}

// Set sets current value of an object.
func (o *Object[T]) Set(value T) {
	o.lock.Lock()
	defer o.lock.Unlock()

	o.state = o.state.update(value)
}

// Subscribe returns a channel that is written to when an object is set to a new value.
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

// Stream returns a stream of values for the given object.
func (o *Object[T]) Stream() Stream[T] {
	return Stream[T]{
		state: o.state,
	}
}
