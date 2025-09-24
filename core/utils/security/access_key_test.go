package security_test

import (
	"testing"
	"time"

	"github.com/cde/go-example/core/utils/security"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccessKey_Encrypt(t *testing.T) {
	tests := []struct {
		name            string
		appKey          string
		timestamp       int64
		expectedEncrypt string
		expectedDecrypt string
	}{
		{
			name:            `TC1. Given app key "secret" and time 1758555936 When encrypt is called Then return QTNpNXA0QitwRElab2Q2dTZWK24rNlBoS09pMC82MG5HRDVYT2Zyam1sQT0=`,
			appKey:          "secret",
			timestamp:       1758555936,
			expectedEncrypt: "QTNpNXA0QitwRElab2Q2dTZWK24rNlBoS09pMC82MG5HRDVYT2Zyam1sQT0=",
			expectedDecrypt: `s:17:"secret@1758555936";`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessKey := security.NewAccessKey(tt.appKey)
			actual := accessKey.Encrypt(time.Unix(tt.timestamp, 0))
			assert.Equal(t, tt.expectedEncrypt, actual)
			decrypt, err := accessKey.Decrypt(actual)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedDecrypt, decrypt)
		})
	}
}
