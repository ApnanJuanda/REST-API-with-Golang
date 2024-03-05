package helper

import (
	"katalisStack.com/practice-golang-restful-api/model/domain"
	"katalisStack.com/practice-golang-restful-api/model/web"
)

func ToCategoryResponse(category *domain.Category) *web.CategoryResponse {
	return &web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
