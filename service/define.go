package main

import (
	"context"
)

// StringService String服务
type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}
