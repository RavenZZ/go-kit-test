package service

import (
	"context"
)

// StringService String服务
type StringService interface {
	Uppercase(ctx context.Context, a string) (str string, err error)
	Count(ctx context.Context, a string) (count int, err error)
	Lowercase(ctx context.Context, a string) (str string, err error)
}
