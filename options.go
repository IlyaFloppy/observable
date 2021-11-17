package observable

type params struct {
	sendCurrent bool
	bufferSize  int
}

type Option func(p params) params

func WithSendCurrent(sendCurrent bool) Option {
	return func(p params) params {
		p.sendCurrent = sendCurrent
		return p
	}
}

func WithBufferSize(bufferSize int) Option {
	return func(p params) params {
		p.bufferSize = bufferSize
		return p
	}
}
