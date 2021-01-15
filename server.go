package todo

import (
	"context"
	"net/http"
	"time"
)

/*Server -*/
type Server struct {
	httpServer *http.Server
}

/*Start -*/
func (s *Server) Start(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

/*Stop -*/
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
