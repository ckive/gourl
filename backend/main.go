package main

import (
	"fmt"

	handler "github.com/ckive/gourl/backend/api"
	"github.com/ckive/gourl/backend/store"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

const ShortLinkLength = 8

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// setting to release mode
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// CORS
	// r.Use(CORSMiddleware())

	// Use the cors middleware with the appropriate configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://gourl.localhost"}, // Replace with your frontend's URL
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Use CORS Default Configs (does not allow all origins)
	// config := cors.DefaultConfig()
	// r.Use(cors.New(config))

	// config := cors.Default() // allows all origins
	// r.Use(config)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note store initialization happens here
	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
