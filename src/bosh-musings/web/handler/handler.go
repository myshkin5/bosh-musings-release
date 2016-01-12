package handler

import (
	"fmt"
	"net/http"
)

type handler struct {
	version string
}

func New(version string) *handler {
	return &handler{
		version: version,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello, world!\nversion = %s", h.version)
}
