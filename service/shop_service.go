package service

import (
	"net/http"

	"github.com/spectertian/dianputest/database"
	"github.com/spectertian/dianputest/model"
)

type ShopRequest struct {
	Name        string `json:"name" binding:"required"`
	Image       string `json:"image"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Status      *int   `json:"status"`
}

type ShopListResult struct {
	List     []model.Shop `json:"list"`
	Total    int64        `json:"total"`
	Page     int          `json:"page"`
	PageSize int          `json:"page_size"`
}

type ShopError struct {
	HTTPCode int
	Message  string
}

func (e *ShopError) Error() string { return e.Message }

func ListShops(page, pageSize int, name string) (*ShopListResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	query := database.DB.Model(&model.Shop{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, &ShopError{HTTPCode: http.StatusInternalServerError, Message: "查询失败"}
	}

	var shops []model.Shop
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&shops).Error; err != nil {
		return nil, &ShopError{HTTPCode: http.StatusInternalServerError, Message: "查询失败"}
	}

	return &ShopListResult{
		List:     shops,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func GetShop(id uint) (*model.Shop, error) {
	var shop model.Shop
	if err := database.DB.First(&shop, id).Error; err != nil {
		return nil, &ShopError{HTTPCode: http.StatusNotFound, Message: "商铺不存在"}
	}
	return &shop, nil
}

func CreateShop(req ShopRequest, userID uint) (*model.Shop, error) {
	status := 1
	if req.Status != nil {
		status = *req.Status
	}
	shop := &model.Shop{
		Name:        req.Name,
		Image:       req.Image,
		Address:     req.Address,
		Phone:       req.Phone,
		Description: req.Description,
		Status:      status,
		UserID:      userID,
	}
	if err := database.DB.Create(shop).Error; err != nil {
		return nil, &ShopError{HTTPCode: http.StatusInternalServerError, Message: "创建商铺失败"}
	}
	return shop, nil
}

func UpdateShop(id uint, req ShopRequest) (*model.Shop, error) {
	var shop model.Shop
	if err := database.DB.First(&shop, id).Error; err != nil {
		return nil, &ShopError{HTTPCode: http.StatusNotFound, Message: "商铺不存在"}
	}

	updates := map[string]interface{}{
		"name":        req.Name,
		"image":       req.Image,
		"address":     req.Address,
		"phone":       req.Phone,
		"description": req.Description,
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if err := database.DB.Model(&shop).Updates(updates).Error; err != nil {
		return nil, &ShopError{HTTPCode: http.StatusInternalServerError, Message: "更新商铺失败"}
	}
	return &shop, nil
}

func DeleteShop(id uint) error {
	var shop model.Shop
	if err := database.DB.First(&shop, id).Error; err != nil {
		return &ShopError{HTTPCode: http.StatusNotFound, Message: "商铺不存在"}
	}
	if err := database.DB.Delete(&shop).Error; err != nil {
		return &ShopError{HTTPCode: http.StatusInternalServerError, Message: "删除商铺失败"}
	}
	return nil
}
