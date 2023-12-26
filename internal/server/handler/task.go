package handler

import (
	"context"
	"fmt"

	"github.com/LOLONEKIT2019/prof_of_work/internal/contract"
	"github.com/LOLONEKIT2019/prof_of_work/internal/services"
	"github.com/LOLONEKIT2019/prof_of_work/pgk/message"
)

type (
	TaskHandler struct {
		taskService services.TaskService
	}
)

func NewTaskHandler() TaskHandler {
	return TaskHandler{
		taskService: services.NewTaskService(),
	}
}

func (s *TaskHandler) RequestTask(ctx context.Context, userInfo string, _ string) (*message.Message, error) {

	task, err := s.taskService.GetTask(ctx, userInfo)
	if err != nil {
		return &message.Message{
			Type: contract.ResponseError,
			Body: fmt.Sprintf("get task: %v", err),
		}, nil
	}
	fmt.Println("get task")

	return &message.Message{
		Type: contract.ResponseTask,
		Body: task,
	}, nil
}

func (s *TaskHandler) RequestQuote(ctx context.Context, userInfo string, body string) (*message.Message, error) {

	quote, err := s.taskService.GetQuote(ctx, userInfo, body)
	if err != nil {
		return &message.Message{
			Type: contract.ResponseError,
			Body: fmt.Sprintf("get quote: %v", err),
		}, nil
	}
	fmt.Println("get quote")

	return &message.Message{
		Type: contract.ResponseQuote,
		Body: quote,
	}, nil
}
