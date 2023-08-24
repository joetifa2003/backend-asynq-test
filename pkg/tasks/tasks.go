package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"golang.org/x/net/context"
)

const (
	TypeHello = "hello"
)

type TaskHelloPayload struct {
	Name string `json:"name"`
}

func NewHelloTask(name string) (*asynq.Task, error) {
	payload, err := json.Marshal(TaskHelloPayload{
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeHello, payload), nil
}

func HandleHelloTask(ctx context.Context, t *asynq.Task) error {
	var payload TaskHelloPayload
	err := json.Unmarshal(t.Payload(), &payload)
	if err != nil {
		return err
	}

	fmt.Println("Hello " + payload.Name)
	return nil
}
