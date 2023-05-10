package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/middleware"
	"goblog/model"
	"goblog/util/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var token string
	var code int
	code = model.CheckLogin(data.Username, data.Password)
	fmt.Println(code)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data.Username,
		"id":      data.ID,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

func LoginFront(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var token string
	var code int
	code = model.CheckLogin(data.Username, data.Password)
	fmt.Println(code)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data.Username,
		"id":      data.ID,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
