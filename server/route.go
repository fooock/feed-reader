package main

import (
	"github.com/gin-gonic/gin"
)

// Route is the struct to create rest resources
type Route struct {
	name    string
	method  string
	pattern string
	handler func(h *gin.Context)
}

// Routes is a set of routes
type Routes []Route

// FeedRoutes for the Feed resource
var FeedRoutes = Routes{
	Route{
		name:    "Get all feeds",
		method:  "GET",
		pattern: "/feeds",
		handler: GetAll,
	},
	Route{
		name:    "Delete a feed",
		method:  "DELETE",
		pattern: "/feed/:id",
		handler: Delete,
	},
	Route{
		name:    "Create a new feed",
		method:  "POST",
		pattern: "/feed",
		handler: Submit,
	},
}
