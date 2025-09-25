package repository

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type BucketS3 struct {
	awsConfig *s3.Client
	bucket    string
}

func (b BucketS3) GeneratePutSignedUrl(ctx context.Context, key string, exp time.Duration) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (b BucketS3) GenerateGetSignedUrl(ctx context.Context, key string, exp time.Duration) (string, error) {
	//TODO implement me
	panic("implement me")
}
