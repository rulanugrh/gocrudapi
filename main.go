package main

import (
	"gocrudapi/configs"
	"gocrudapi/entity"
	"gocrudapi/handler"
	"gocrudapi/repository"
	"gocrudapi/route"
	services "gocrudapi/service"
	"log"
)

func main() {
	log.Println("Server running at localhost:3000")
	db := configs.GetConnection()
	db.AutoMigrate(&entity.Book{}, &entity.Category{})

	bookRepository := repository.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	route.Run(bookHandler)
}
