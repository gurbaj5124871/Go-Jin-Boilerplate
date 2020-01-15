package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// InitiliseRoutes func to initilise all the routes and export router type
func InitiliseRoutes() *gin.Engine {
	router = gin.Default()

	// loading static files
	router.Static("/public", "./public")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.LoadHTMLGlob("views/*")

	// rendering index and not found routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Go Gin Server",
		})
	})
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"title":         "Go Gin Server",
			"error_title":   "Not Found",
			"error_message": "Page Not Found",
		})
	})

	return router
}
