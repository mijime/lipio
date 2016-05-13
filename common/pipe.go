package common

import "io"

type PipeFactory interface {
	Match(o Option) bool
	Create(o Option) (Pipe, error)
}

type Pipe interface {
	Execute(w io.Writer, r io.Reader) (int64, error)
}

func NewPipe(o Option) (Pipe, error) {
	for _, f := range pipeFactories {
		if f.Match(o) {
			return f.Create(o)
		}
	}

	return nil, NotMatchError
}

func AddPipeFactory(f PipeFactory) {
	pipeFactories = append(pipeFactories, f)
}

var pipeFactories = make([]PipeFactory, 0)
