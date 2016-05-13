package serv

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/mijime/lipio/common"
)

func init() {
	common.AddHandlerFactory(httpHandlerFactory{})
}

var httpRegexp = regexp.MustCompile("https?")

type httpHandlerFactory struct{}

func (f httpHandlerFactory) Create(o common.Option) (http.Handler, error) {
	return &httpHandler{}, nil
}

func (f httpHandlerFactory) Match(o common.Option) bool {
	return httpRegexp.Match([]byte(o.URL.Scheme))
}

type httpHandler struct{}

func (s *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mr, readErr := r.MultipartReader()

	if readErr != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, readErr)
		log.Println(readErr)
		return
	}

	for part, eof := mr.NextPart(); eof == nil; part, eof = mr.NextPart() {
		_, copyErr := io.Copy(os.Stdout, part)

		if copyErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, copyErr)
			log.Println(copyErr)
			return
		}
	}
}
