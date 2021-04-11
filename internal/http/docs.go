package http

import (
	"net/http"
)

type docsEndpoint struct {
	html []byte
}

func (d docsEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	if _, err := w.Write(d.html); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
