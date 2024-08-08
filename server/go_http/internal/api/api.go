package api

import "net/http"

func NewApi() *http.Server {
	srv := newServer()

	return &http.Server{
		Addr:    ":8080",
		Handler: srv,
	}
}
