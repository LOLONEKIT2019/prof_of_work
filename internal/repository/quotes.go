package repository

import (
	"context"
	"math/rand"
)

var quotesArray = []string{
	"quote 1",
	"quote 2",
	"quote 3",
	"quote 4",
}

type (
	QuoteRepository interface {
		GetRandomQuote(ctx context.Context) (string, error)
	}

	quotes struct{}
)

func NewQuotesRepository() QuoteRepository {
	return &quotes{}
}

func (r *quotes) GetRandomQuote(_ context.Context) (string, error) {
	return quotesArray[rand.Intn(len(quotesArray))], nil
}
