package pkg

import "github.com/gin-gonic/gin"

func (s *Server) routes() *gin.Engine {
	r := s.Router

	v1 := r.Group("/v1/api")
	{
		v1.GET("/status", s.ApiStatus())
		v1.POST("/mail", s.TestMail())
	}

	return r
}
