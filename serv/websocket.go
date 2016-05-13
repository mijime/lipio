package serv

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/mijime/lipio/common"
	"golang.org/x/net/websocket"
)

func init() {
	common.AddHandlerFactory(websocketHandlerFactory{})
}

var websocketRegexp = regexp.MustCompile("wss?")

type websocketHandlerFactory struct{}

func (f websocketHandlerFactory) Create(o common.Option) (http.Handler, error) {
	return websocket.Handler(stdoutHandler), nil
}

func (f websocketHandlerFactory) Match(o common.Option) bool {
	return websocketRegexp.Match([]byte(o.URL.Scheme))
}

func stdoutHandler(c *websocket.Conn) {
	_, err := io.Copy(os.Stdout, c)

	if err != nil {
		log.Println(err)
	}
}
