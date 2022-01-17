package observable

type params struct {
	sendCurrent bool
	bufferSize  int
}

// Option is an Option for Object.Subscribe method.
type Option func(p params) params

// WithSendCurrent makes Subscribe put current Object value into the returned channel when sendCurrent is true.
// Only new values that are set after Subscribe will be sent over channel otherwise.
func WithSendCurrent(sendCurrent bool) Option {
	return func(p params) params {
		p.sendCurrent = sendCurrent
		return p
	}
}

// WithBufferSize specifies channel size.
func WithBufferSize(bufferSize int) Option {
	return func(p params) params {
		p.bufferSize = bufferSize
		return p
	}
}
