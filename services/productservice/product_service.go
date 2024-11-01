package productservice

import (
	"errors"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/models"
	"github.com/agussuartawan/golang-pos/repositories/productrepository"
	"github.com/agussuartawan/golang-pos/repositories/unitrepository"
)

func Create(req request.ProductRequest) (*uint, error) {
	// pastikan unitId ada di db
	exists, err := unitrepository.IsExists(req.UnitId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("unitId not found")
	}

	// map req to model for product
	productModel := models.Product{
		UnitId: req.UnitId,
		Name:   req.Name,
		Size:   req.Size,
	}
	if err := productrepository.CreateProduct(&productModel); err != nil {
		return nil, err
	}

	// map req to model for price
	priceModel := models.ProductPrice{
		ProductID: productModel.ID,
		Value:     req.Price.Value,
		StartDate: req.Price.StartDate,
		EndDate:   req.Price.EndDate,
	}
	if err := productrepository.CreateProductPrice(priceModel); err != nil {
		return nil, err
	}

	return &productModel.ID, nil
}
