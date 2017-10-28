package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	version = "v0.1.0"
	header  = "Feed Reader server " + version
)

// Options to start the server
type Options struct {
	port         string
	host         string
	databaseFile string
}

func main() {
	fmt.Println(header)

	// parse command line options
	port := flag.String("p", "8080", "Port where the server will be listening")
	host := flag.String("x", "0.0.0.0", "Host to attach")

	// start in debug mode?
	debug := flag.Bool("dev", false, "Use this flag to start the server in debug mode")
	// database file
	databaseFile := flag.String("d", "/tmp/feed-server.db", "Sqlite database file")
	flag.Parse()

	log.Printf("Start server in %s:%s\n", *host, *port)

	// change debug flag in gin
	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}
	options := &Options{
		port:         *port,
		host:         *host,
		databaseFile: *databaseFile,
	}
	// create and start the server instance
	server, err := Create(*options)
	if err != nil {
		log.Fatalln(err)
	}
	server.Start()

	// release
	defer server.Close()
}
