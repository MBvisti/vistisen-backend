package app

func (s *Server) routes() {
	// Health check
	s.router.HandleFunc("/v1/api/test", s.ApiStatus())
}
