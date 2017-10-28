package main

import "github.com/gin-gonic/gin"
import "fmt"

// Server is the main entry point of the program
type Server struct {
	router *gin.Engine
}

// Create the server instance
func Create() *Server {
	router := gin.Default()
	server := &Server{router: router}
	return server
}

// Start the server with the given options
func (s *Server) Start(options Options) {
	s.router.Run(fmt.Sprintf("%v:%v", options.host, options.port))
}
