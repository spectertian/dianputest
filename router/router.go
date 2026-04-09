package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spectertian/dianputest/controller"
	"github.com/spectertian/dianputest/middleware"
)

func Setup(r *gin.Engine) {
	api := r.Group("/api")

	// Public routes
	api.POST("/register", controller.Register)
	api.POST("/login", controller.Login)

	// Protected routes
	auth := api.Group("/")
	auth.Use(middleware.JWTAuth())
	{
		auth.POST("/logout", controller.Logout)

		auth.GET("/shops", controller.ListShops)
		auth.GET("/shops/:id", controller.GetShop)
		auth.POST("/shops", controller.CreateShop)
		auth.PUT("/shops/:id", controller.UpdateShop)
		auth.DELETE("/shops/:id", controller.DeleteShop)
	}
}
