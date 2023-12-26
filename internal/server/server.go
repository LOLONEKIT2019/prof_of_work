package server

import (
	"context"
	"fmt"

	"github.com/LOLONEKIT2019/prof_of_work/internal/config"
	"github.com/LOLONEKIT2019/prof_of_work/internal/contract"
	"github.com/LOLONEKIT2019/prof_of_work/internal/server/handler"
	"github.com/LOLONEKIT2019/prof_of_work/pgk/message"
	serverLib "github.com/LOLONEKIT2019/prof_of_work/pgk/server"
)

func Start() error {

	taskHandler := handler.NewTaskHandler()

	handlerFn := func(ctx context.Context, userInfo string, messageType message.Type, body string) (*message.Message, error) {
		switch messageType {
		case contract.RequestTask:
			return taskHandler.RequestTask(ctx, userInfo, body)
		case contract.RequestQuote:
			return taskHandler.RequestQuote(ctx, userInfo, body)
		}
		return nil, fmt.Errorf("unsupported message type")
	}

	err := serverLib.RunServer(config.Config.ServerPort, handlerFn)
	if err != nil {
		return fmt.Errorf("run server: %w", err)
	}

	return nil
}
