package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spectertian/dianputest/service"
	"github.com/spectertian/dianputest/utils"
)

func Register(c *gin.Context) {
	var req service.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	user, err := service.Register(req)
	if err != nil {
		if authErr, ok := err.(*service.AuthError); ok {
			utils.Fail(c, authErr.HTTPCode, authErr.Message)
		} else {
			utils.Fail(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.Success(c, "注册成功", gin.H{
		"id":       user.ID,
		"username": user.Username,
		"nickname": user.Nickname,
	})
}

func Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	user, err := service.Login(req)
	if err != nil {
		utils.Fail(c, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "Token 生成失败")
		return
	}

	utils.Success(c, "登录成功", gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"nickname": user.Nickname,
		},
	})
}

func Logout(c *gin.Context) {
	utils.Success(c, "登出成功", nil)
}
