package main

import (
	"log"
	"scaleX/internal/handlers/restHandler"
	"scaleX/internal/repository"
	"scaleX/internal/usecase"
)

func main() {

	userRepo := repository.NewUserRepo()
	authService := usecase.NewAuthService(userRepo)
	bookService := usecase.NewBookService(userRepo)
	openApiService := usecase.NewOpenApiService()

	handler := restHandler.NewRestHandler(authService, bookService, openApiService)

	echoServer := restHandler.Echo(handler)

	log.Fatal("failed to start server", echoServer.Start(":8888"))

}
