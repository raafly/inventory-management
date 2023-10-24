package main

import (
	"net/http"
	
	"github.com/go-playground/validator/v10"
	"github.com/raafly/inventory-management/config"
	"github.com/raafly/inventory-management/controller"
	"github.com/raafly/inventory-management/repository"
	"github.com/raafly/inventory-management/route"
	"github.com/raafly/inventory-management/service"
)

func main() {
	DB := config.NewDB()
	Validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, DB, Validate)
	userController := controller.NewUserController(userService)

	itemRepository := repository.NewItemRepository()
	itemService := service.NewItemService(itemRepository, DB, Validate)
	itemController := controller.NewItemController(itemService)

	router := route.NewRouter(userController, itemController)

	server := http.Server {
		Addr: "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}