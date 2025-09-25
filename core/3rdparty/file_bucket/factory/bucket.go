package factory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	appConfig "github.com/cde/go-example/config"
	"github.com/cde/go-example/core/3rdparty/file_bucket/repository"
)

func MakeBucketS3(cfg *appConfig.Config) *repository.BucketS3 {
	awsConfig, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(cfg.AwsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AwsAccessKeyId,
			cfg.AwsSecretAccessKey,
			cfg.AwsSessionToken,
		)),
	)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	client := s3.NewFromConfig(awsConfig)
	presigner := s3.NewPresignClient(client)

	return repository.NewBucketS3(presigner, cfg.AwsS3Bucket)
}

func ResolveBucketRepository(cfg *appConfig.Config) repository.BucketInterface {
	return MakeBucketS3(cfg)
}
