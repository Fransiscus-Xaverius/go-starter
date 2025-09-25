package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/cde/go-example/core/security/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccessKey_Encrypt(t *testing.T) {
	tests := []struct {
		ctx             context.Context
		name            string
		appKey          string
		timestamp       int64
		expectedEncrypt string
		expectedDecrypt string
	}{
		{
			name:            `TC1. Given app key "secret" and time 1758555936 When encrypt is called Then return QTNpNXA0QitwRElab2Q2dTZWK24rNlBoS09pMC82MG5HRDVYT2Zyam1sQT0=`,
			ctx:             context.Background(),
			appKey:          "secret",
			timestamp:       1758555936,
			expectedEncrypt: "QTNpNXA0QitwRElab2Q2dTZWK24rNlBoS09pMC82MG5HRDVYT2Zyam1sQT0=",
			expectedDecrypt: `s:17:"secret@1758555936";`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessKey := repository.NewAccessKeyAesSha256Cbc(tt.appKey)
			actual := accessKey.Encrypt(tt.ctx, time.Unix(tt.timestamp, 0))
			assert.Equal(t, tt.expectedEncrypt, actual)
			decrypt, err := accessKey.Decrypt(tt.ctx, actual)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedDecrypt, decrypt)
		})
	}
}
