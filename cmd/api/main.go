package main

import (
	"log"
	"scaleX/internal/handlers/restHandler"
	"scaleX/internal/usecase"
)

func main() {

	authService := usecase.NewAuthService()
	bookService := usecase.NewBookService()

	handler := restHandler.NewRestHandler(authService, bookService)

	echoServer := restHandler.Echo(handler)

	log.Fatal("failed to start server", echoServer.Start(":8888"))

}
