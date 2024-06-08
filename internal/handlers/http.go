package handlers

import (
	"fmt"
	"log"
	"net/http"
	"pavelkononov/resume/public"
)

type HttpHandler interface {
}

func MakeHttpServer() {
	h := http.NewServeMux()
	frontend, er1 := public.GetFS()
	if er1 != nil {
		log.Fatal(er1)
	}
	fileServer := http.FileServer(http.FS(frontend))
	h.Handle("GET /", fileServer)

	server := &http.Server{Addr: ":8080", Handler: h}
	fmt.Println("Started on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
