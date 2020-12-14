package http

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type Server struct {
	engine *gin.Engine
}

type Route struct {
	Method       string
	Path         string
	IsAuthorized bool
	Function     gin.HandlerFunc
}

type Option struct {
	Address string
	User    string
	Pass    string
}

func NewHTTPServer() *Server {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	s := Server{
		engine: engine,
	}
	return &s
}

func (s *Server) Start(address string) {
	go func(address string) {
		log.Infof("Start listening HTTP on address: %v", address)
		err := s.engine.Run(address)
		if err != nil {
			log.Fatal(err)
		}
	}(address)
}

func (s *Server) AddRoutes(routes ...Route) {
	for _, r := range routes {
		s.engine.Handle(r.Method, r.Path, r.Function)
	}
}
