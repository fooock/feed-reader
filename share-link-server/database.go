package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Database is the struct to access to data
type Database struct {
	db *gorm.DB
}

// Start is the function that create the database schema
func (d *Database) Start() {
	if !d.db.HasTable(&Feed{}) {
		d.db.CreateTable(&Feed{})
	}
}

// Close the database connection
func (d *Database) Close() {
	d.db.Close()
}
