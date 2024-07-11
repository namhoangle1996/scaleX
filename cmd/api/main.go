package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"scaleX/config"
	"scaleX/internal/handlers/restHandler"
	"scaleX/internal/repository"
	"scaleX/internal/usecase"
	"time"
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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := echoServer.Start(configs.Application.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			echoServer.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := echoServer.Shutdown(ctx); err != nil {
		echoServer.Logger.Fatal(err)
	}

}
