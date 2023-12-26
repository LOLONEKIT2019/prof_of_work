package services

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/LOLONEKIT2019/prof_of_work/internal/config"
	"github.com/LOLONEKIT2019/prof_of_work/internal/pow"
	"github.com/LOLONEKIT2019/prof_of_work/internal/repository"
)

type (
	TaskService interface {
		GetTask(ctx context.Context, userInfo string) (string, error)
		GetQuote(ctx context.Context, userInfo, hash string) (string, error)
	}

	taskService struct {
		attemptRepository repository.AttemptRepository
		quoteRepository   repository.QuoteRepository
		timeNow           func() time.Time
	}
)

func NewTaskService() TaskService {
	return &taskService{
		attemptRepository: repository.NewAttemptRepository(),
		quoteRepository:   repository.NewQuotesRepository(),
		timeNow: func() time.Time {
			return time.Now()
		},
	}
}

func (s *taskService) GetTask(ctx context.Context, userInfo string) (string, error) {

	err := s.attemptRepository.Save(ctx, userInfo, s.timeNow().Add(config.Config.TaskTTL))
	if err != nil {
		return "", fmt.Errorf("save attempt: %w", err)
	}

	return strconv.Itoa(rand.Intn(9999)), nil
}

func (s *taskService) GetQuote(ctx context.Context, userInfo, hash string) (string, error) {

	expiration, err := s.attemptRepository.GetExpiration(ctx, userInfo)
	if err != nil {
		return "", fmt.Errorf("get attempt expiration: %w", err)
	}

	if s.timeNow().After(expiration) {
		return "", errors.New("task expired")
	}

	if !pow.IsHashCorrect(hash, config.Config.ZerosCount) {
		return "", errors.New("validation failed")
	}

	err = s.attemptRepository.Remove(ctx, userInfo)
	if err != nil {
		return "", fmt.Errorf("remove attempt: %w", err)
	}

	quote, err := s.quoteRepository.GetRandomQuote(ctx)
	if err != nil {
		return "", fmt.Errorf("get quote: %w", err)
	}

	return quote, nil
}
