package repository

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type BucketS3 struct {
	presignClient *s3.PresignClient
	bucket        string
}

func NewBucketS3(presignClient *s3.PresignClient, bucket string) *BucketS3 {
	return &BucketS3{presignClient: presignClient, bucket: bucket}
}

func (b BucketS3) GeneratePutSignedUrl(ctx context.Context, key string, exp time.Duration) (string, error) {
	req, err := b.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(b.bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(exp))
	if err != nil {
		return "", err
	}
	return req.URL, nil

}

func (b BucketS3) GenerateGetSignedUrl(ctx context.Context, key string, exp time.Duration) (string, error) {
	req, err := b.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(b.bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(exp))
	if err != nil {
		return "", err
	}
	return req.URL, nil

}
