package http

import (
	"booking-insights/internal/infrastructure/config"
	"context"
	"fmt"
	"net"
	"net/http"
)

// NewServer creates a new  http serrver with the confg passed in
func NewServer(ctx context.Context, cfg config.Config) *http.Server {
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.HttPort),
		ReadTimeout:    cfg.HttpReadTimeout,
		WriteTimeout:   cfg.HttpWriteTimeout,
		MaxHeaderBytes: cfg.HttpMaxHeaderBytes,
		BaseContext: func(n net.Listener) context.Context {
			return ctx
		},

		// TODO (revisit if got time)
		// ConnContext: func(connCtx context.Context, c net.Conn) context.Context{
		// 	return ctx
		// },
	}

	statsendp := StatsHandler{}
	http.Handle("/stats", &statsendp)

	return server
}
