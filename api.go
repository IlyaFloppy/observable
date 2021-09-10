package observable

func New[T any](value T) *Object[T] {
	return &Object[T]{state: newState(value)}
}
