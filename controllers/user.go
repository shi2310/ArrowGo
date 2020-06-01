package controllers

import (
	"ArrowGo/middleware"
	"ArrowGo/models"
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register 注册
func Register(c *gin.Context) {
	result := models.ResponseData{
		Success: true,
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		result.Success = false
		result.Msg = "解析失败"
		c.JSON(http.StatusUnauthorized, result)
		return
	}
	if _, err := models.GetUserByUserName(user.UserName); err == nil {
		result.Success = false
		result.Msg = "用户已存在"
		c.JSON(http.StatusBadRequest, result)
		return
	}
	if err := models.AddUser(&user); err != nil {
		result.Success = false
		result.Msg = "数据库异常"
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	c.JSON(http.StatusOK, result)
}

// Login 登录
func Login(c *gin.Context) {
	result := models.ResponseData{
		Success: true,
	}
	username := c.PostForm("username")
	pwd := c.PostForm("password")
	if len(username) == 0 || len(pwd) == 0 {
		result.Success = false
		result.Msg = "账号或密码不能为空"
		c.JSON(http.StatusOK, result)
	} else {
		user, err := models.GetUserByUserName(username)
		if err != nil {
			result.Success = false
			result.Msg = "用户不存在"
			c.JSON(http.StatusOK, result)
			return
		}
		password := fmt.Sprintf("%x", md5.Sum([]byte(user.UserName+pwd)))
		if user.Pwd != password {
			result.Success = false
			result.Msg = "密码错误"
			c.JSON(http.StatusOK, result)
			return
		}
		// 登录成功，生成Token
		token, err := middleware.GenToken(user.UserName)
		if err != nil {
			result.Success = false
			result.Msg = "Token生成失败"
			c.JSON(http.StatusOK, result)
			return
		}
		result.Data = token
		c.JSON(http.StatusOK, result)
	}
}

// AddUser 用户添加
func AddUser(c *gin.Context) {
	var user models.User
	result := models.ResponseData{
		Success: true,
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		result.Success = false
		result.Msg = "解析失败"
		c.JSON(http.StatusUnauthorized, result)
		return
	}
	if _, err := models.GetUserByUserName(user.UserName); err == nil {
		result.Success = false
		result.Msg = "用户已存在"
		c.JSON(http.StatusOK, result)
		return
	}
	if err := models.AddUser(&user); err != nil {
		result.Success = false
		result.Msg = "数据库异常"
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	c.JSON(http.StatusOK, result)
}
