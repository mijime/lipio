package serv

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	AddHandlerFactory(httpHandlerFactory{})
}

type httpHandlerFactory struct{}

func (f httpHandlerFactory) Create(o Option) (http.Handler, error) {
	return &httpHandler{}, nil
}

func (f httpHandlerFactory) Match(o Option) bool {
	return true
}

type httpHandler struct{}

func (s *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mr, readErr := r.MultipartReader()

	if readErr != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, readErr)
		return
	}

	for part, eof := mr.NextPart(); eof == nil; part, eof = mr.NextPart() {
		_, copyErr := io.Copy(os.Stdout, part)

		if copyErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, copyErr)
			return
		}
	}
}
