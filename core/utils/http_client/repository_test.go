package http_client_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	presentationDto "github.com/cde/go-example/core/presentation/dto"
	"github.com/cde/go-example/core/utils/http_client"
	"github.com/cde/go-example/src/modules/user/dto"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationGetUser(t *testing.T) {
	// Create context
	ctx := context.Background()

	// Use the real HttpClientRepository
	httpClientRepo := http_client.NewHttpClientRepository(&http.Client{}).EnableDebug()

	// Use v2.Get to make a real HTTP GET request to the Cat Facts API
	response, err := http_client.Get[presentationDto.Response[dto.UserResponse]](ctx, httpClientRepo, "http://localhost:3000/users/1", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	actualResponse, err := json.Marshal(response.Content)
	assert.NoError(t, err)

	expectedResponse := `{"status":true,"data":{"id":1,"name":"John Doe","email":"john.doe@example.com"}}`
	assert.Equal(t, expectedResponse, string(actualResponse))
}
