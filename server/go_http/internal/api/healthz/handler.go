package healthz

import (
	"io"
	"net/http"
)

func HandleHealthz() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Healthz")
		},
	)
}
