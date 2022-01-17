package observable

// Stream is an entity that allows to monitor Object changes without spawning a goroutine unlike Subscribe.
type Stream[T any] struct {
	state *state[T]
}

// Value returns current value.
func (s *Stream[T]) Value() T {
	return s.state.value
}

// Next advances stream. Should not be called unless Changes returns a closed channel.
func (s *Stream[T]) Next() T {
	s.state = s.state.next

	return s.Value()
}

// Changes returns a channel that is closed when next value is available.
func (s *Stream[T]) Changes() <-chan struct{} {
	return s.state.done
}

// WaitNext waits for next value.
func (s *Stream[T]) WaitNext() {
	<-s.Changes()
}

// HasNext returns true if stream has next value.
func (s *Stream[T]) HasNext() bool {
	select {
	case <-s.state.done:
		return true
	default:
		return false
	}
}
