package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("Port not provided. Add `PORT=8080` to .env")
	}
	server := &http.Server{Addr: fmt.Sprintf("%s:%s", "", port), Handler: h}
	fmt.Println("Started on http://localhost:" + port)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
