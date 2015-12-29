package handler

import (
	"fmt"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello, world!")
}
