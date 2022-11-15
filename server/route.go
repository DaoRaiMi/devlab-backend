package server

func (s *Server) registerRoute() {
	userV1 := s.engin.Group("/api/v1/user/")
	userV1.POST("login", s.login)
}
