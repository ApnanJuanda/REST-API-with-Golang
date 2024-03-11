package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"katalisStack.com/practice-golang-restful-api/app"
	"katalisStack.com/practice-golang-restful-api/controller"
	"katalisStack.com/practice-golang-restful-api/helper"
	"katalisStack.com/practice-golang-restful-api/repository"
	"katalisStack.com/practice-golang-restful-api/service"
	"net/http"
	"os"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	// mapping
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	fmt.Println("My Application is running")
	err := godotenv.Load("config/.env")
	helper.PanicIfError(err)
	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router,
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
