package service

import (
	"errors"
	"net/http"

	"github.com/spectertian/dianputest/database"
	"github.com/spectertian/dianputest/model"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthError struct {
	HTTPCode int
	Message  string
}

func (e *AuthError) Error() string { return e.Message }

func Register(req RegisterRequest) (*model.User, error) {
	var existing model.User
	if err := database.DB.Where("username = ?", req.Username).First(&existing).Error; err == nil {
		return nil, &AuthError{HTTPCode: http.StatusConflict, Message: "用户名已存在"}
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, &AuthError{HTTPCode: http.StatusInternalServerError, Message: "密码加密失败"}
	}

	user := &model.User{
		Username: req.Username,
		Password: string(hashed),
		Nickname: req.Nickname,
	}
	if err := database.DB.Create(user).Error; err != nil {
		return nil, &AuthError{HTTPCode: http.StatusInternalServerError, Message: "注册失败"}
	}
	return user, nil
}

func Login(req LoginRequest) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	return &user, nil
}
