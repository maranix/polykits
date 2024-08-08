package routes

import (
	"net/http"

	"github.com/maranix/polykits/server/go_http/internal/api/healthz"
)

func AddRoutes(mux *http.ServeMux) {
	mux.Handle("/healthz", healthz.HandleHealthz())
}
