package main

import (
	"gofirebase/api"
	"gofirebase/config"
	"gofirebase/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize new gin engine (for server)
	r := gin.Default()

	// create/configure database instance
	db := config.CreateDatabase()

	// configure firebase
	firebaseAuth := config.SetupFirebase()

	// set db & firebase to gin context with a middleware to all incoming request
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", firebaseAuth)
	})

	// using the auth middleware to validate api request
	r.Use(middleware.AuthMiddleware)

	// Routes definition for finding and creating artist
	r.GET("/artist", api.FindArtists)
	r.POST("/artist", api.CreateArtist)

	// start server
	r.Run(":5000")
}
