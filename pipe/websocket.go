package pipe

import (
	"io"
	"log"
	"regexp"

	"github.com/mijime/lipio/common"
	"golang.org/x/net/websocket"
)

func init() {
	common.AddPipeFactory(websocketPipeFactory{})
}

type websocketPipeFactory struct{}

var websocketRegexp = regexp.MustCompile("wss?")

func (r websocketPipeFactory) Match(o common.Option) bool {
	return websocketRegexp.Match([]byte(o.URL.Scheme))
}

func (r websocketPipeFactory) Create(o common.Option) (common.Pipe, error) {
	return &websocketPipe{Url: o.Raw, Origin: o.URL.Host}, nil
}

type websocketPipe struct {
	Url, Origin string
}

func (p *websocketPipe) Execute(w io.Writer, r io.Reader) (int64, error) {
	ws, err := websocket.Dial(p.Url, "", p.Origin)

	if err != nil {
		return 0, err
	}

	defer ws.Close()
	go func() {
		_, err := io.Copy(w, ws)

		if err != nil {
			log.Println(err)
		}
	}()
	return io.Copy(ws, r)
}
