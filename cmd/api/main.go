package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joetifa2003/confessions/cmd/api/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	controllers := []controllers.Controller{
		controllers.AsynqmonHandler{},
		controllers.Hello{},
	}
	for _, controller := range controllers {
		controller.Mount(e)
	}

	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
