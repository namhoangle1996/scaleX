package main

import (
	"log"
	"scaleX/config"
	"scaleX/internal/handlers/restHandler"
	"scaleX/internal/repository"
	"scaleX/internal/usecase"
)

func main() {

	configs, err := config.GetConfig()
	if err != nil {
		panic("Can not get config ,error " + err.Error())
	}

	userRepo := repository.NewUserRepo()
	authService := usecase.NewAuthService(userRepo)
	bookService := usecase.NewBookService(userRepo)
	openApiService := usecase.NewOpenApiService(configs)

	handler := restHandler.NewRestHandler(authService, bookService, openApiService)

	echoServer := restHandler.Echo(handler)

	log.Fatal("failed to start server", echoServer.Start(configs.Application.Port))

}
