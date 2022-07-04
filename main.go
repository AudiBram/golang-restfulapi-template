package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"golang-restfulapi/app"
	controller2 "golang-restfulapi/controller"
	"golang-restfulapi/helper"
	"golang-restfulapi/middleware"
	"golang-restfulapi/repository"
	service2 "golang-restfulapi/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service2.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller2.NewController(categoryService)

	r := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(r),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
