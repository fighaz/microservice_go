package main

import (
	"fmt"
	"net/http"
	"user/config"
	"user/controller"
	"user/repository"
	"user/usecase"
)

func main() {
	db := config.Connect()

	userrepo := repository.NewUserRepository(db)
	userusecase := usecase.NewUserUsecase(userrepo)
	controller.UserRouter(userusecase)
	fmt.Println("server started at localhost:6000")
	http.ListenAndServe(":6000", nil)
}

// $2a$14$bWGHM7MnhHw/WgFHAtocpu3RrlR4V73Icws4akmgV6rV85Kpmoyfu
