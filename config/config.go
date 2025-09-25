package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName    string `envconfig:"APP_NAME" default:"demo-api"`
	AppVersion string `envconfig:"APP_VERSION" default:"v1.0.0"`
	AppEnv     string `envconfig:"APP_ENV" default:"development"`
	AppPort    int    `envconfig:"APP_PORT" default:"3000"`
	AppDebug   bool   `envconfig:"APP_DEBUG" default:"true"`

	MySQLHost              string `envconfig:"MYSQL_HOST" default:"localhost"` // host.docker.internal
	MySQLPort              int    `envconfig:"MYSQL_PORT" default:"3306"`
	MySQLUsername          string `envconfig:"MYSQL_USERNAME" default:"root"`
	MySQLPassword          string `envconfig:"MYSQL_PASSWORD" default:"password12345"`
	MySQLDbName            string `envconfig:"MYSQL_DBNAME" default:"demo"`
	MySQLMaxIdleConnection int    `envconfig:"MYSQL_MAX_IDLE_CONNECTION" default:"10"`
	MySQLMaxOpenConnection int    `envconfig:"MYSQL_MAX_OPEN_CONNECTION" default:"100"`
	MySQLConnMaxLifetime   int    `envconfig:"MYSQL_CONN_MAX_LIFETIME" default:"5"` // in minutes

	CorsAllowOrigins string `envconfig:"CORS_ALLOW_ORIGINS" default:"*"` // "https://gofiber.io, https://gofiber.net"
	AllowHeaders     string `envconfig:"CORS_ALLOW_HEADERS" default:"Origin, Content-Type, Accept, X-Requested-Id"`

	AwsAccessKeyId     string `envconfig:"AWS_ACCESS_KEY_ID" default:""`
	AwsSecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY" default:""`
	AwsSessionToken    string `envconfig:"AWS_SESSION_TOKEN" default:""`
	AwsRegion          string `envconfig:"AWS_REGION" default:"ap-southeast-1"`
	AwsS3Bucket        string `envconfig:"AWS_S3_BUCKET" default:""`

	AyoAuthTimeout int `envconfig:"AYO_AUTH_TIMEOUT" default:"5"` // in seconds
}

var cfg *Config

func Get() *Config {
	if cfg == nil {
		cfg = &Config{}
		envconfig.MustProcess("", cfg)
	}

	return cfg
}
