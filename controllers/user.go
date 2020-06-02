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
		result.Msg = "解析失败" + err.Error()
		c.JSON(http.StatusUnauthorized, result)
		return
	}
	if _, err := models.GetUserByUserName(user.UserName); err == nil {
		result.Success = false
		result.Msg = "用户已存在"
		c.JSON(http.StatusBadRequest, result)
		return
	}
	user.Pwd = fmt.Sprintf("%x", md5.Sum([]byte(user.UserName+user.Pwd)))
	if err := models.AddUser(&user); err != nil {
		result.Success = false
		result.Msg = "数据库异常" + err.Error()
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
	var model models.User
	if err := c.ShouldBindJSON(&model); err != nil {
		result.Success = false
		result.Msg = "解析失败" + err.Error()
		c.JSON(http.StatusUnauthorized, result)
		return
	}
	if len(model.UserName) == 0 || len(model.Pwd) == 0 {
		result.Success = false
		result.Msg = "账号或密码不能为空"
		c.JSON(http.StatusOK, result)
	} else {
		user, err := models.GetUserByUserName(model.UserName)
		if err != nil {
			result.Success = false
			result.Msg = "用户不存在" + err.Error()
			c.JSON(http.StatusOK, result)
			return
		}
		password := fmt.Sprintf("%x", md5.Sum([]byte(user.UserName+model.Pwd)))
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
