package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/VsevolodCher/eventhub/internal/config"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)

	cfg := config.Load()
	slog.Info("event-service starting", "addr", cfg.HTTPAddr)

	if err := http.ListenAndServe(cfg.HTTPAddr, mux); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
