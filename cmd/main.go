package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/raafly/inventory-management/internal/listing"
	"github.com/raafly/inventory-management/pkg/config"
)

func main() {
	DB := config.NewDB()
	Validate := validator.New()

	userRepository := listing.NewUserRepository(DB)
	userService := listing.NewUserService(userRepository, DB, Validate)
	userHandler := listing.NewUserController(userService)

	itemRepository := listing.NewItemRepository(DB)
	itemService := listing.NewItemService(itemRepository, DB, Validate)
	itemHandler := listing.NewItemController(itemService)

	categoryRepository := listing.NewCategoryRepository()
	categoryService := listing.NewCategoryService(categoryRepository, DB, Validate)
	categoryHandler := listing.NewCategoryHandler(categoryService)

	router := listing.NewRouter(userHandler, itemHandler, categoryHandler)

	server := http.Server {
		Addr: "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}