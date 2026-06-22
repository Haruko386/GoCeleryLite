package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Server struct {
	addr string
	http *http.Server
}

func New(addr string) *Server {
	s := &Server{
		addr: addr,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", s.handleHealth)

	s.http = &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s
}

func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.http.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return s.http.Shutdown(shutdownCtx)
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}

// handleHealth
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJson(w, http.StatusOK, map[string]any{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

// writeJson
func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}
