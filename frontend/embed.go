package frontend

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist/*
var content embed.FS

func HandleUi() {
	dist, err := fs.Sub(content, "dist")
	if err != nil {
		panic(err)
	}

	withResponseHeaders := func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Set some header.
			w.Header().Add("Access-Control-Allow-Origin", "*")
			// Serve with the actual handler.
			h.ServeHTTP(w, r)
		}
	}

	http.Handle("/", withResponseHeaders(http.FileServer(http.FS(dist))))
}
