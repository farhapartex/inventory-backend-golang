package server

import (
	"github.com/gin-gonic/gin"
	"github.com/goupp-backend/controller"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Check system health
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm healthy",
		})
	})

	// API endpoints
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			auth := v1.Group("auth")
			{
				auth.POST("/register", controller.Register)

				// auth.GET("/test", func(ctx *gin.Context) {
				// 	ctx.JSON(200, gin.H{
				// 		"message": "I'm test healthy",
				// 	})
				// })
			}

			users := v1.Group("users")
			{
				users.GET("/", func(ctx *gin.Context) {
					ctx.JSON(200, gin.H{
						"message": "I'm healthy",
					})
				})

				users.GET("/test", func(ctx *gin.Context) {
					ctx.JSON(200, gin.H{
						"message": "I'm test healthy",
					})
				})
			}
		}
	}

	return router
}