package main

import (
	"gulbeyaz-backend/internal/auth"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Handler interface {
	RegisterRoutes() (string, http.Handler)
}

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8081",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	authHandler := &auth.Handler{}

	handlers := []Handler{
		authHandler,
	}

	for _, handler := range handlers {
		path, handler := handler.RegisterRoutes()
		mux.Handle(path, handler)
	}

	log.Println("server listening on: 8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
