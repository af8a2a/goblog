package routes

import (
	"github.com/gin-gonic/gin"
	"goblog/util"
	"net/http"
)

func InitRouter() {
	gin.SetMode(util.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
	r.Run(util.HttpPort)
}
