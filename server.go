package innotaxiuser

import "net/http"

//Server ...
type Server struct {
	httpServer *http.Server
}

//Run : method that start our server, that will work until error returns
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    port,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}
