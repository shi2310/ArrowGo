package controllers

import (
	"ArrowGo/dto"
	"ArrowGo/middleware"
	"ArrowGo/models"
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 注册
// @Description 注册接口
// @Accept  json
// @Produce  json
// @Param model body dto.UserDTO true "用户注册模型"
// @Success 200 {object} models.ResponseData {"success":true,"data":{},"msg":null}
// @Router /account/register [post]
func Register(c *gin.Context) {
	result := models.ResponseData{
		Success: true,
	}
	var model dto.UserDTO
	if err := c.ShouldBindJSON(&model); err != nil {
		result.Success = false
		result.Msg = "解析失败" + err.Error()
		c.JSON(http.StatusUnauthorized, result)
		return
	}
	if _, err := models.GetUserByUserName(model.UserName); err == nil {
		result.Success = false
		result.Msg = "用户已存在"
		c.JSON(http.StatusBadRequest, result)
		return
	}

	var user models.User
	user.UserName = model.UserName
	user.Pwd = fmt.Sprintf("%x", md5.Sum([]byte(model.UserName+model.Pwd)))
	if err := models.AddUser(&user); err != nil {
		result.Success = false
		result.Msg = "数据库异常" + err.Error()
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary 登录
// @Accept  json
// @Produce  json
// @Param model body dto.LoginDTO true "登录模型"
// @Success 200 {object} models.ResponseData {"success":true,"data":{},"msg":null}
// @Router /account/login [post]
func Login(c *gin.Context) {
	result := models.ResponseData{
		Success: true,
	}
	var model dto.LoginDTO
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

// @Summary 添加用户
// @Accept multipart/form-data
// @Produce  json
// @Param Authorization	header string true "Toke:格式如Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param username formData string true "用户名"
// @Param pwd formData string true "密码"
// @Param photo formData file false "图像"
// @Success 200 {object} models.ResponseData {"success":true,"data":{},"msg":null}
// @Router /user/add [post]
func AddUser(c *gin.Context) {
	var user models.User
	result := models.ResponseData{
		Success: true,
	}
	username := c.PostForm("username")
	if _, err := models.GetUserByUserName(username); err == nil {
		result.Success = false
		result.Msg = "用户已存在"
		c.JSON(http.StatusOK, result)
		return
	}

	user.UserName = username
	user.Pwd = fmt.Sprintf("%x", md5.Sum([]byte(username+c.PostForm("pwd"))))
	file, ex := FormUpload(c, "photo")
	if ex == nil {
		user.Photo = file.FileMd5
	}

	if err := models.AddUser(&user); err != nil {
		result.Success = false
		result.Msg = "数据库异常"
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary 修改个人密码
// @Accept  json
// @Produce  json
// @Param Authorization	header string true "Toke:格式如Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param model body dto.ChangePwdDTO true "新密码"
// @Success 200 {object} models.ResponseData {"success":true,"data":{},"msg":null}
// @Router /user/changePwd [PUT]
func ChangePwd(c *gin.Context) {
	result := models.ResponseData{
		Success: true,
	}
	userName := c.MustGet("username").(string)
	user, err := models.GetUserByUserName(userName)
	if err != nil {
		result.Success = false
		result.Msg = "用户不存在"
		c.JSON(http.StatusUnauthorized, result)
		return
	}

	var model dto.ChangePwdDTO
	if err := c.ShouldBindJSON(&model); err != nil {
		result.Success = false
		result.Msg = "解析失败" + err.Error()
		c.JSON(http.StatusBadRequest, result)
		return
	}

	user.Pwd = fmt.Sprintf("%x", md5.Sum([]byte(user.UserName+model.Pwd)))
	if err := models.ChangePwd(user); err != nil {
		result.Success = false
		result.Msg = "数据库异常"
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	c.JSON(http.StatusOK, result)
}
