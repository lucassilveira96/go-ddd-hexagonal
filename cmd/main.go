package main

import (
	"github.com/joho/godotenv"
	"go-ddd-hexagonal/app/service"
	"go-ddd-hexagonal/infrastructure/api"
	"go-ddd-hexagonal/infrastructure/database"
	"go-ddd-hexagonal/infrastructure/repository"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro carregando arquivo .env")
	}

	db := database.InitDB()

	productRepo := repository.NewSQLProductRepository(db)
	productService := service.NewProductService(productRepo)

	router := http.NewServeMux()
	api.RegisterProductRoutes(router, productService)

	http.ListenAndServe(":5050", router)
}
