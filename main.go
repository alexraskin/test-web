package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Request received", "method", r.Method, "url", r.URL.String(), "ip", r.RemoteAddr)
		w.Write([]byte("Hello, World!"))
	})
	http.ListenAndServe(":8080", r)
}
