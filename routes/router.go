package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web_app/controllers"
	"web_app/logger"
	"web_app/middlewares"
	"web_app/pkg/snowflake"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")

	v1.GET("/", func(c *gin.Context) {
		id := snowflake.GenID()
		c.String(http.StatusOK, strconv.FormatInt(id, 10))
	})

	v1.POST("/signup", controllers.SignUpHandler)
	v1.POST("/login", controllers.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		v1.GET("/posts", controllers.GetPostListHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
