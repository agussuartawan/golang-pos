package productrepository

import (
	"github.com/agussuartawan/golang-pos/core/config"
	"github.com/agussuartawan/golang-pos/data/request"
	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/agussuartawan/golang-pos/models"
	"time"
)

func CreateProduct(model *models.Product) error {
	return config.DB.Create(&model).Error
}

func List(data *[]response.ProductResponse, param *request.ProductParam) error {
	type ProductResponse struct {
		ID       uint    `json:"id"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Size     uint    `json:"size"`
		UnitId   uint    `json:"unitId"`
		UnitName string  `json:"unitName"`
	}
	var tempData []ProductResponse

	// Subquery untuk mengambil harga terkini yang berlaku hari ini dari tabel `product_prices`
	priceSubQuery := config.DB.
		Select("value").
		Table("product_prices").
		Where("product_prices.product_id = products.id").
		Where("start_date <= ? AND end_date >= ?", time.Now(), time.Now()).
		Order("start_date DESC").
		Limit(1)

	// Query utama untuk mengambil data produk dengan harga terkini
	query := config.DB.Model(&models.Product{}).
		Select("products.id, products.name, products.size, units.id as unit_id, units.name as unit_name, (?) as price", priceSubQuery).
		Joins("join units on units.id = products.unit_id")

	if param.Name != nil {
		query = query.Where("products.name ILIKE ?", "%"+*param.Name+"%")
	}

	if param.Query != nil {
		query = query.Where("name ilike '%?%' and size = ?", *param.Query, *param.Query)
	}

	query = param.Paginate(query)
	switch param.SortBy {
	case "name":
		query = query.Order("products.name " + param.SortBy)
	case "size":
		query = query.Order("products.size " + param.SortBy)
	default:
		query = query.Order("products.created_at desc")
	}

	if err := query.Find(&tempData).Error; err != nil {
		return err
	}

	// Mapping ke ProductResponse
	var responseData []response.ProductResponse
	for _, result := range tempData {
		responseData = append(responseData, response.ProductResponse{
			ID:    result.ID,
			Name:  result.Name,
			Price: result.Price,
			Size:  result.Size,
			Unit: response.Unit{
				ID:   result.UnitId,
				Name: result.UnitName,
			},
		})
	}
	*data = responseData

	return nil
}
