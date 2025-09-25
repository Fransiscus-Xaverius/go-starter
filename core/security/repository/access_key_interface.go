package repository

import (
	"context"
	"time"
)

//go:generate mockgen -destination=mocks/mock_AccessKeyInterface.go -package=mocks . AccessKeyInterface
type (
	AccessKeyInterface interface {
		Encrypt(ctx context.Context, currTime time.Time) string
		Decrypt(ctx context.Context, encrypted string) (string, error)
	}
)