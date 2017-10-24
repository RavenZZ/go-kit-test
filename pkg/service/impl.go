package service

import (
	"context"
	"errors"
	"strings"
)

type stringService struct{}

func NewStringService() StringService {
	return stringService{}
}

func (stringService) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(_ context.Context, s string) (int, error) {
	return len(s), nil
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")
