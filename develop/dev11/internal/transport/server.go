package transport

import (
	"net/http"
	"time"
)

type Server struct {
	serv *http.Server
}

func (s *Server) RunServer(port string, mux http.Handler) error {
	s.serv = &http.Server{
		Addr:           ":" + port,
		Handler:        mux,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.serv.ListenAndServe()
}
