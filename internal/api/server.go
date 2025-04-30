package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
    router *gin.Engine
}

func NewServer() *Server {
    r := gin.Default()
    return &Server{router: r}
}

func (s *Server) InitializeRoutes() {
    // Initialize your routes here
    // Example: s.router.GET("/brands", s.GetBrands)
}

func (s *Server) Run(addr string) error {
    return s.router.Run(addr)
}

func (s *Server) HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "up"})
}