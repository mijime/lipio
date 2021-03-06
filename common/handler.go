package common

import "net/http"

func NewHandler(o Option) (http.Handler, error) {
	for _, f := range handlerFactories {
		if f.Match(o) {
			return f.Create(o)
		}
	}

	return nil, NotMatchError
}

type HandlerFactory interface {
	Create(o Option) (http.Handler, error)
	Match(o Option) bool
}

func AddHandlerFactory(f HandlerFactory) {
	handlerFactories = append(handlerFactories, f)
}

var handlerFactories = make([]HandlerFactory, 0)
