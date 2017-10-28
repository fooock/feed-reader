package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Feed is a resource of interest
type Feed struct {
	Model
	Name   string `gorm:"not null;unique" json:"name"`
	Author string `gorm:"not null" json:"author"`
}

// Model for each feed
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// Feeds is a set of feeds
type Feeds []Feed

// FeedHandler is the struct to call all feed handlers
type FeedHandler struct {
	database *Database
}

// GetAll is the function to retrieve all feeds
func (f *FeedHandler) GetAll(c *gin.Context) {
	var feeds Feeds
	f.database.db.Find(&feeds)
	if len(feeds) == 0 {
		c.JSON(404, gin.H{
			"status":  404,
			"message": "Feeds not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"feeds": feeds,
	})
}

// Delete the given feed
func (f *FeedHandler) Delete(c *gin.Context) {

}

// Submit create a new feed
func (f *FeedHandler) Submit(c *gin.Context) {
	// get the data from the post. Note that the user can be annonymous
	author := c.DefaultPostForm("author", "annon")
	feed := c.PostForm("feed")

	if len(feed) == 0 {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Can't be empty feed",
		})
		return
	}
	saveFeed := &Feed{
		Name:   feed,
		Author: author,
	}
	if count := f.database.db.Create(saveFeed).RowsAffected; count != 1 {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Feed can't be saved",
		})
	} else {
		// feed created
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Feed processed",
		})
	}
}
