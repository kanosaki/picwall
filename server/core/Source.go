package core

type StreamControlMessage int

const (
	STREAM_STOP StreamControlMessage = 0
	STREAM_TICK                      = -1
)

func (scm StreamControlMessage) IsDrainRequest() bool {
	return scm > 0
}

func (scm StreamControlMessage) AsDrainRequest() int {
	if scm.IsDrainRequest() {
		return int(scm)
	} else {
		panic("Invalid operation! check IsDrainRequest first")
	}
}

func (scm StreamControlMessage) IsSpecial() bool {
	return scm <= 0
}

type ControlChannel chan StreamControlMessage

type Source interface {
	Faucet() <-chan *Entry
	Drain(count int)
	Shutdown()
}
