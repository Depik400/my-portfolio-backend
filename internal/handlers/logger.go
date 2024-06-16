package handlers

import (
	"fmt"
	"net/http"
)

type HTTPHandler func(w http.ResponseWriter, req *http.Request)

func PublicHttpLoggerFactory(f HTTPHandler) HTTPHandler {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			fmt.Printf("address: %s, remote: %s\n", req.Host, req.RemoteAddr)
		}
		f(w, req)
	}
}
