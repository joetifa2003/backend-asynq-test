package controllers

import (
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/joetifa2003/confessions/pkg/tasks"
	"github.com/labstack/echo/v4"
)

type Hello struct {
}

func (h Hello) Mount(e *echo.Echo) {
	e.GET("/hello", h.Get)
}

func (h Hello) Get(c echo.Context) error {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     ":6379",
		Password: "devpassword",
	})

	helloTask, err := tasks.NewHelloTask("from handler")
	if err != nil {
		return err
	}

	_, err = client.Enqueue(helloTask)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"hello": "world",
	})
}
