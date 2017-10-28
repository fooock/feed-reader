package main

import "github.com/gin-gonic/gin"
import "fmt"

const APIV1 = "api/v1"

// Server is the main entry point of the program
type Server struct {
	router *gin.Engine
}

// Create the server instance
func Create() *Server {
	router := gin.Default()
	// feed endpoints
	feedEndpoints(router)
	server := &Server{router: router}
	return server
}

// Start the server with the given options
func (s *Server) Start(options Options) {
	s.router.Run(fmt.Sprintf("%v:%v", options.host, options.port))
}

// feedEndpoints register all feed endpoints with the given engine
func feedEndpoints(e *gin.Engine) {
	for _, route := range FeedRoutes {
		group := e.Group(APIV1)

		// register get
		if route.method == "GET" {
			group.GET(route.pattern, route.handler)
		}
		// register delete
		if route.method == "DELETE" {
			group.DELETE(route.pattern, route.handler)
		}
	}
}
