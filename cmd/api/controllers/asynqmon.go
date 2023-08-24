package controllers

import (
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/labstack/echo/v4"
)

type AsynqmonHandler struct{}

func (am AsynqmonHandler) Mount(e *echo.Echo) {
	g := e.Group("/processor")

	mon := asynqmon.New(asynqmon.Options{
		RootPath: "/processor",
		RedisConnOpt: asynq.RedisClientOpt{
			Addr:     ":6379",
			Password: "devpassword",
		},
	})

	g.Any("/*", echo.WrapHandler(mon))
}
