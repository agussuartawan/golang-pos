package productrepository

import (
	"github.com/agussuartawan/golang-pos/core/config"
	"github.com/agussuartawan/golang-pos/models"
)

func CreateProductPrice(model models.ProductPrice) error {
	return config.DB.Create(&model).Error
}
