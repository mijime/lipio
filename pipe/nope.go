package pipe

import (
	"io"
	"regexp"
)

func init() {
	AddPipeFactory(nopePipeFactory{})
}

type nopePipeFactory struct{}

var nopeRegexp = regexp.MustCompile("nope")

func (f nopePipeFactory) Match(o Option) bool {
	return nopeRegexp.Match([]byte(o.Scheme))
}

func (f nopePipeFactory) Create(o Option) (Pipe, error) {
	return &nopePipe{}, nil
}

type nopePipe struct {
}

func (p *nopePipe) Execute(w io.Writer, r io.Reader) (int64, error) {
	return io.Copy(w, r)
}
