package main

import (
	"golang-restapi/app"
	"golang-restapi/controller"
	"golang-restapi/exception"
	"golang-restapi/helper"
	"golang-restapi/middleware"
	"golang-restapi/repository"
	"golang-restapi/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewcategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    ":3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
