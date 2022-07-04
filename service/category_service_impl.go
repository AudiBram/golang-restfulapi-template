package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang-restfulapi/exception"
	"golang-restfulapi/helper"
	"golang-restfulapi/model/domain"
	"golang-restfulapi/model/web"
	"golang-restfulapi/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, DB: DB, Validate: validate}
}

func (s *CategoryServiceImpl) Create(c context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Id:   0,
		Name: request.Name,
	}
	category = s.CategoryRepository.Save(c, tx, category)
	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) Update(c context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.CategoryRepository.FindById(c, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	category.Name = request.Name

	category = s.CategoryRepository.Update(c, tx, category)
	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) Delete(c context.Context, categoryId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.CategoryRepository.FindById(c, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	s.CategoryRepository.Delete(c, tx, category)
}

func (s *CategoryServiceImpl) FindById(c context.Context, categoryId int) web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.CategoryRepository.FindById(c, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) FindAll(c context.Context) []web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := s.CategoryRepository.FindAll(c, tx)
	return helper.ToCategoryResponses(categories)
}
