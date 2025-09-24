package repository

import (
	"time"
)

//go:generate mockgen -destination=mocks/mock_AccessKey.go -package=mocks . AccessKey
type (
	AccessKey interface {
		Encrypt(currTime time.Time) string
		Decrypt(encrypted string) (string, error)
	}
)