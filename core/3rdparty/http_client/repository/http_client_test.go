package repository_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cde/go-example/core/3rdparty/http_client"
	"github.com/cde/go-example/core/3rdparty/http_client/repository"
	presentationDto "github.com/cde/go-example/core/presentation/dto"
	"github.com/cde/go-example/src/modules/user/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationGetUser(t *testing.T) {
	// Create context
	ctx := context.Background()

	// Use the real HttpClientInterface
	httpClientRepo := repository.NewHttpClientRepository(&http.Client{
		Timeout: 5 * time.Second,
	}).EnableDebug()

	// Use v2.Get to make a real HTTP GET request to the Cat Facts API
	response, err := http_client.Get[presentationDto.Response[dto.UserResponse]](ctx, httpClientRepo, "http://localhost:3000/users/1", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	actualResponse, err := json.Marshal(response.Content)
	assert.NoError(t, err)

	expectedResponse := `{"status":true,"data":{"id":1,"name":"John Doe","email":"john.doe@example.com"}}`
	assert.Equal(t, expectedResponse, string(actualResponse))
}

func TestIntegrationCreateUser(t *testing.T) {
	tests := []struct {
		name               string
		request            *dto.UserRequest
		expectedStatusCode int
		expectedResponse   string
		expectedError      string
	}{
		{
			name:               `TC1. Given empty struct When call Post Create User Then return 400 bad request`,
			request:            &dto.UserRequest{},
			expectedStatusCode: fiber.StatusBadRequest,
			expectedError:      `ERRDEMO400002`,
		},
		{
			name: `TC2. Given valid struct When call Post Create User Then return 201 created`,
			request: &dto.UserRequest{
				Name:     "John Doe",
				Email:    fmt.Sprintf("%s-johndone@gmail.com", utils.UUIDv4()),
				Password: "secret",
			},
			expectedStatusCode: fiber.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create context
			ctx := context.Background()

			// Use the real HttpClientInterface
			httpClientRepo := repository.NewHttpClientRepository(&http.Client{
				Timeout: 5 * time.Second,
			}).EnableDebug()

			// Use v2.Get to make a real HTTP GET request to the Cat Facts API
			response, err := http_client.Post[*dto.UserRequest, presentationDto.Response[dto.UserResponse]](
				ctx,
				httpClientRepo,
				"http://localhost:3000/users",
				tt.request,
				map[string]string{
					"content-type": "application/json",
				},
			)
			if tt.expectedError == "" {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedStatusCode, response.StatusCode)
				log.Printf("%+v", response.Content)
				return
			}

			actualErrorStr := err.Error()
			log.Println(actualErrorStr)
			assert.Contains(t, actualErrorStr, tt.expectedError)
		})
	}
}
