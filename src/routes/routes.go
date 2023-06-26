package routes

import (
	"dumblink-be-go/src/controllers/auth"
	"dumblink-be-go/src/controllers/link"
	"dumblink-be-go/src/controllers/profile"
	authMiddleware "dumblink-be-go/src/middlewares/auth"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Router() *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.POST("/api/v1/login", auth.Login)
	router.POST("/api/v1/register", auth.Register)

	router.GET("/api/v1/link/:id", link.GetLink)
	router.GET("/api/v1/link/:id/count", link.CountLInk)
	router.GET("/api/v1/link-uniqid/:uniqid", link.GetLinkByUniqid)

	router.Use(authMiddleware.Auth())
	{
		router.GET("/api/v1/profile", profile.GetProfile)
		router.PUT("/api/v1/profile", profile.UpdateProfile)
		router.DELETE("/api/v1/profile", profile.DeleteProfile)

		router.POST("/api/v1/link", link.AddLink)
		router.GET("/api/v1/link", link.GetLinks)
		router.PUT("/api/v1/link", link.UpdateLink)
		router.DELETE("/api/v1/link", link.DeleteLink)
	}

	return router
}
