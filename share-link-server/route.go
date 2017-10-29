package main

// Route is the struct to create rest resources
type Route struct {
	name    string
	method  string
	pattern string
}

// Routes is a set of routes
type Routes []Route

// FeedRoutes for the Feed resource
var FeedRoutes = Routes{
	Route{
		name:    "Get all feeds",
		method:  "GET",
		pattern: "/feeds",
	},
	Route{
		name:    "Delete a feed",
		method:  "DELETE",
		pattern: "/feed/:id",
	},
	Route{
		name:    "Create a new feed",
		method:  "POST",
		pattern: "/feed",
	},
}
