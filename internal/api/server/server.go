package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinasadeghi83/ghashoghchi/internal/api/v1/user"
	"gorm.io/gorm"
)

type Server struct {
	Engine *gin.Engine
	Addr   string
	DB     *gorm.DB
}

func NewServer(addr string, db *gorm.DB) *Server {
	return &Server{
		Engine: gin.Default(),
		Addr:   addr,
		DB:     db,
	}
}

func (s *Server) SetupRoutes() {
	apiV1 := s.Engine.Group("/api/v1")

	//Register modules
	user.RegisterRoutes(apiV1.Group("/auth"), s.DB)

	//Health Check
	s.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
}

func (s *Server) Start() error {
	log.Printf("Server listening on %s", s.Addr)
	return s.Engine.Run(s.Addr)
}
