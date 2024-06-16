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
	h.HandleFunc("GET /", PublicHttpLoggerFactory(fileServer.ServeHTTP))
	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("Port not provided. Add `PORT=8080` to .env")
	}

	host, exists := os.LookupEnv("HOST")
	if !exists {
		log.Fatal("Host not provided. Add `PORT=8080` to .env")
	}

	server := &http.Server{Addr: fmt.Sprintf("%s:%s", host, port), Handler: h}
	fmt.Printf("Started on http://%s:%s\n", host, port)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
