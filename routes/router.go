package routes

import (
	"github.com/gin-gonic/gin"
	v1 "goblog/api/v1"
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
		// 用户模块的路由接口

		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		// 文章模块的路由接口
		router.POST("category/add", v1.AddCategory)
		router.GET("category", v1.GetCate)
		router.PUT("category/:id", v1.EditCate)
		router.DELETE("category/:id", v1.DeleteCate)
	}

	r.Run(util.HttpPort)
}
