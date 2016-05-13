package pipe

import (
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
)

func init() {
	AddPipeFactory(httpPipeFactory{})
}

type httpPipeFactory struct{}

var httpRegexp = regexp.MustCompile("https?://.+")

func (r httpPipeFactory) Match(o Option) bool {
	return httpRegexp.Match([]byte(o.Scheme))
}

func (r httpPipeFactory) Create(o Option) (Pipe, error) {
	return &httpPipe{Url: o.Scheme, Method: "POST"}, nil
}

type httpPipe struct {
	Url, Method, Type string
}

func (p *httpPipe) Execute(w io.Writer, r io.Reader) (int64, error) {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	errHandler := make(chan error)

	go func() {
		part, partErr := mw.CreateFormField("data")
		if partErr != nil {
			mw.Close()
			pw.Close()
			errHandler <- partErr
			return
		}

		_, copyErr := io.Copy(part, r)
		if copyErr != nil {
			mw.Close()
			pw.Close()
			errHandler <- copyErr
			return
		}

		mw.Close()
		pw.Close()
		errHandler <- nil
	}()

	req, reqErr := http.NewRequest(p.Method, p.Url, pr)
	if reqErr != nil {
		return 0, reqErr
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())

	client := &http.Client{}
	res, resErr := client.Do(req)
	if resErr != nil {
		return 0, resErr
	}
	defer res.Body.Close()

	mwErr := <-errHandler
	close(errHandler)

	if mwErr != nil {
		return 0, mwErr
	}

	return io.Copy(w, res.Body)
}
