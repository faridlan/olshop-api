package model

import (
	"github.com/faridlan/olshop-rest-api/model/domain"
	"github.com/faridlan/olshop-rest-api/model/web/categoryweb"
)

func ToCategoryResponse(category domain.Category) categoryweb.CategoryResponse {
	return categoryweb.CategoryResponse{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}

func ToCategoryResponses(categories []domain.Category) []categoryweb.CategoryResponse {
	categoryResponses := []categoryweb.CategoryResponse{}
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
