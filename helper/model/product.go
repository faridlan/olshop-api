package model

import (
	"github.com/faridlan/olshop-rest-api/model/domain"
	"github.com/faridlan/olshop-rest-api/model/web/productweb"
)

func ToProductResponse(product domain.Product) productweb.ProductResponse {
	productResponse := productweb.ProductResponse{
		Id:         product.Id,
		Name:       product.Name,
		Price:      product.Price,
		Quantity:   product.Quantity,
		CategoryId: product.CategoryId,
		Desription: product.Desription,
		ImageUrl:   product.ImageUrl,
		CreatedAt:  product.CreatedAt,
	}

	return productResponse
}

func ToProductResponses(products []domain.Product) []productweb.ProductResponse {
	productResponses := []productweb.ProductResponse{}

	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}

	return productResponses
}
