package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web_app/controllers"
	"web_app/logger"
	"web_app/pkg/snowflake"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		id := snowflake.GenID()
		c.String(http.StatusOK, strconv.FormatInt(id, 10))
	})

	r.POST("/signup", controllers.SignUpHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
