package app

import (
	"github.com/julienschmidt/httprouter"
	"golang-restfulapi/controller"
	"golang-restfulapi/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	r := httprouter.New()

	r.GET("/api/categories", categoryController.FindAll)
	r.GET("/api/categories/:categoryId", categoryController.FindById)
	r.PUT("/api/categories/:categoryId", categoryController.Update)
	r.POST("/api/categories", categoryController.Create)
	r.DELETE("/api/categories/:categoryId", categoryController.Delete)

	r.PanicHandler = exception.ErrorHandler

	return r
}
