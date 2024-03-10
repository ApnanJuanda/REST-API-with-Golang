package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"katalisStack.com/practice-golang-restful-api/helper"
	"katalisStack.com/practice-golang-restful-api/model/domain"
	"katalisStack.com/practice-golang-restful-api/model/web"
	"katalisStack.com/practice-golang-restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service CategoryServiceImpl) Create(ctx context.Context, request *web.CategoryCreateRequest) *web.CategoryResponse {
	// validate request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// db begin
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}
	// call repository
	category = *service.CategoryRepository.Save(ctx, tx, &category)
	return helper.ToCategoryResponse(&category)
}

func (service CategoryServiceImpl) Update(ctx context.Context, request *web.CategoryUpdateRequest) *web.CategoryResponse {
	// validate request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// db begin
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// findById
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name
	// update
	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	// db begin
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// findById
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	// delete
	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId int) *web.CategoryResponse {
	// db begin
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// findById
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []*web.CategoryResponse {
	// db begin
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// findAll
	categories := service.CategoryRepository.FindAll(ctx, tx)

	var categoryResponses []*web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}
	return categoryResponses
}
