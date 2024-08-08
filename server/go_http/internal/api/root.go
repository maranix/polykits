package api

import (
	"net/http"

	"github.com/maranix/polykits/server/go_http/internal/routes"
)

func newServer() http.Handler {
	mux := http.NewServeMux()

	routes.AddRoutes(mux)

	var handler http.Handler = mux
	return handler
}
