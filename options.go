package observable

type params struct {
	withCurrent bool
	bufferSize  int
}

// Option is an Option for Object.Subscribe method.
type Option func(p params) params

// WithCurrent makes Subscribe put current Object value into the returned channel when sendCurrent is true.
// Only new values that are set after Subscribe will be sent over channel otherwise.
func WithCurrent(withCurrent bool) Option {
	return func(p params) params {
		p.withCurrent = withCurrent
		return p
	}
}

// WithBuffer specifies channel buffer size.
func WithBuffer(bufferSize int) Option {
	return func(p params) params {
		p.bufferSize = bufferSize
		return p
	}
}
