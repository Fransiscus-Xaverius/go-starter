package repository

import (
	"context"
	"time"
)

//go:generate mockgen -destination=mocks/mock_BucketInterface.go -package=mocks . BucketInterface
type BucketInterface interface {
	GeneratePutSignedUrl(ctx context.Context, key string, exp time.Duration) (string, error)
	GenerateGetSignedUrl(ctx context.Context, key string, exp time.Duration) (string, error)
}
