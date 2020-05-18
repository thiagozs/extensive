package server

import (
	"extensive-number/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Engine *gin.Engine
}

func New(port string, debug bool) *Server {
	// Set are debug or not
	gin.SetMode(gin.ReleaseMode)
	if debug {
		gin.SetMode(gin.DebugMode)
	}

	// load gin framework
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return &Server{port, r}
}

// Run server
func (s *Server) Run() error {
	return s.Engine.Run(s.Port)
}

// RegisterRoutes bind the route to engine
func (s *Server) RegisterRoutes() {
	s.Engine.GET("/:number", s.GetFromAPI())
	s.Engine.GET("/:number/ping", s.GetPong())
}

// GetFromAPI handler req to api
func (s *Server) GetFromAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		number := c.Param("number")
		ex := lib.NewExtensive()
		str := ex.Convert(number)

		if str == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid content"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"extenso": str})
	}
}

// GetPong healthcheck api
func (s *Server) GetPong() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "pong")
	}
}
