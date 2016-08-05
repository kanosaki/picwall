package core

type Stream chan *Entry

type ReadStream <-chan *Entry

type Source interface {
	Faucet() ReadStream
	Drain(count int)
	Shutdown()
}
