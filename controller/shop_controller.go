package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spectertian/dianputest/service"
	"github.com/spectertian/dianputest/utils"
)

func ListShops(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.Query("name")

	result, err := service.ListShops(page, pageSize, name)
	if err != nil {
		if shopErr, ok := err.(*service.ShopError); ok {
			utils.Fail(c, shopErr.HTTPCode, shopErr.Message)
		} else {
			utils.Fail(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.Success(c, "success", result)
}

func GetShop(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "无效的商铺 ID")
		return
	}

	shop, err := service.GetShop(uint(id))
	if err != nil {
		if shopErr, ok := err.(*service.ShopError); ok {
			utils.Fail(c, shopErr.HTTPCode, shopErr.Message)
		} else {
			utils.Fail(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.Success(c, "success", shop)
}

func CreateShop(c *gin.Context) {
	var req service.ShopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	userID, _ := c.Get("userID")
	shop, err := service.CreateShop(req, userID.(uint))
	if err != nil {
		if shopErr, ok := err.(*service.ShopError); ok {
			utils.Fail(c, shopErr.HTTPCode, shopErr.Message)
		} else {
			utils.Fail(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.Success(c, "创建成功", shop)
}

func UpdateShop(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "无效的商铺 ID")
		return
	}

	var req service.ShopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	shop, err := service.UpdateShop(uint(id), req)
	if err != nil {
		if shopErr, ok := err.(*service.ShopError); ok {
			utils.Fail(c, shopErr.HTTPCode, shopErr.Message)
		} else {
			utils.Fail(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.Success(c, "更新成功", shop)
}

func DeleteShop(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "无效的商铺 ID")
		return
	}

	if err := service.DeleteShop(uint(id)); err != nil {
		if shopErr, ok := err.(*service.ShopError); ok {
			utils.Fail(c, shopErr.HTTPCode, shopErr.Message)
		} else {
			utils.Fail(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.Success(c, "删除成功", nil)
}
