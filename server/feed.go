package main

import (
	"github.com/gin-gonic/gin"
)

// Feed is a resource of interest
type Feed struct {
	name string
}

// FeedHandler is the struct to call all feed handlers
type FeedHandler struct {
	database *Database
}

// GetAll is the function to retrieve all feeds
func (f *FeedHandler) GetAll(c *gin.Context) {

}

// Delete the given feed
func (f *FeedHandler) Delete(c *gin.Context) {

}

// Submit create a new feed
func (f *FeedHandler) Submit(c *gin.Context) {

}
