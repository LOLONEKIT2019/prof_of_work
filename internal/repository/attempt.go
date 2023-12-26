package repository

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type (
	AttemptRepository interface {
		Save(ctx context.Context, id string, ttl time.Time) error
		GetExpiration(ctx context.Context, id string) (time.Time, error)
		Remove(ctx context.Context, id string) error
	}

	attemptRepository struct {
		data map[string]time.Time // id:expiration
	}
)

func NewAttemptRepository() AttemptRepository {
	return &attemptRepository{data: make(map[string]time.Time)}
}

func (r *attemptRepository) Save(_ context.Context, id string, ttl time.Time) error {
	_, found := r.data[id]
	if found {
		return errors.New("attempt already exist")
	}
	r.data[id] = ttl
	return nil
}

func (r *attemptRepository) GetExpiration(_ context.Context, id string) (time.Time, error) {
	ttl, found := r.data[id]
	if !found {
		return time.Time{}, fmt.Errorf("attempt with id %s not found", id)
	}
	return ttl, nil
}

func (r *attemptRepository) Remove(_ context.Context, id string) error {
	delete(r.data, id)
	return nil
}
