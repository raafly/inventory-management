package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/raafly/inventory-management/internal/listing"
	"github.com/raafly/inventory-management/pkg/config"
	"github.com/raafly/inventory-management/pkg/route"
)

func main() {
	DB := config.NewDB()
	Validate := validator.New()

	userRepository := listing.NewUserRepository(DB)
	userService := listing.NewUserService(userRepository, DB,Validate)
	userHandler := listing.NewUserController(userService)

	itemRepository := listing.NewItemRepository(DB)
	itemService := listing.NewItemService(itemRepository, DB, Validate)
	itemHandler := listing.NewItemController(itemService)

	categoryRepository := listing.NewCategoryRepository(DB)
	categoryService := listing.NewCategoryService(categoryRepository, Validate)
	categoryHandler := listing.NewCategoryHandler(categoryService)

	historyRepository := listing.NewHistoryRepository(DB)
	historyService := listing.NewHistoryService(historyRepository)
	historyHandler := listing.NewHistoryHandler(historyService)

	router := route.NewRouter(userHandler, itemHandler, categoryHandler, historyHandler)
  aurhMiddleware := middleware.NewAuthMiddleware(router)

	server := http.Server {
		Addr: "localhost:3000",
		Handler: aurhMiddleware,
	}

	server.ListenAndServe()
}
