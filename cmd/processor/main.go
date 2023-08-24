package main

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/joetifa2003/confessions/pkg/tasks"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     redisAddr,
			Password: "devpassword",
		},
		asynq.Config{},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeHello, tasks.HandleHelloTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
