package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "Port number for the server")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		b, _ := w.Write([]byte(`Hello from simple microservice written in Go!`))
		log := slog.With(
			slog.String("ip", r.RemoteAddr),
			slog.String("method", r.Method),
			slog.Int("bytes", b),
			slog.Int("status", http.StatusOK),
		)
		log.Info("incoming request")
	})

	slog.Info("starting server", slog.String("port", *port))
	_ = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", *port), mux)
}
