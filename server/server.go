package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const apiV1 = "api/v1"

// Server is the main entry point of the program
type Server struct {
	options  Options
	router   *gin.Engine
	database *Database
}

// Create the server instance using the given options
func Create(opts Options) (*Server, error) {
	// default engine
	router := gin.Default()

	// database
	db, err := gorm.Open("sqlite3", opts.databaseFile)
	if err != nil {
		return nil, err
	}
	database := &Database{db}

	// create the server struct with all required data
	server := &Server{
		options:  opts,
		router:   router,
		database: database,
	}
	// feed endpoints
	feedEndpoints(router, database)

	return server, nil
}

// Start the server with the given options
func (s *Server) Start() {
	s.router.Run(fmt.Sprintf("%v:%v", s.options.host, s.options.port))
}

// Close the server connection and all of its resources
func (s *Server) Close() {
	s.database.Close()
}

// feedEndpoints register all feed endpoints with the given engine
func feedEndpoints(e *gin.Engine, d *Database) {
	feedHandler := &FeedHandler{database: d}
	for _, route := range FeedRoutes {
		group := e.Group(apiV1)

		// register get
		if route.method == "GET" {
			group.GET(route.pattern, feedHandler.GetAll)
		}
		// register delete
		if route.method == "DELETE" {
			group.DELETE(route.pattern, feedHandler.Delete)
		}
		// submit new feed
		if route.method == "POST" {
			group.POST(route.pattern, feedHandler.Submit)
		}
	}
}
