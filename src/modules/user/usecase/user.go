package usecase

import (
	"context"
	"time"

	appContext "github.com/cde/go-example/core/context"
	appError "github.com/cde/go-example/core/error"
	"github.com/cde/go-example/src/modules/user/dto"
	"github.com/cde/go-example/src/modules/user/entity"
	userRepository "github.com/cde/go-example/src/modules/user/repository"
)

type userUseCase struct {
	userRepository userRepository.UserInterface
}

func NewUserUseCase(userRepository userRepository.UserInterface) UserUseCaseInterface {
	return &userUseCase{userRepository: userRepository}
}

func (u userUseCase) CreateUser(ctx context.Context, request *dto.UserRequest) (*dto.UserResponse, error) {
	var (
		logger = appContext.LoggerFromContext(ctx)
	)

	logger.Info("userUseCase.CreateUser")
	nowUnix := time.Now().Unix()
	newRecord, err := u.userRepository.Create(
		ctx,
		&entity.User{
			Name:      request.Name,
			Email:     request.Email,
			Password:  request.Password,
			CreatedAt: nowUnix,
			CreatedBy: 1,
			UpdatedAt: nowUnix,
			UpdatedBy: 1,
		},
	)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    newRecord.ID,
		Name:  newRecord.Name,
		Email: newRecord.Email,
	}, err
}

func (u userUseCase) GetUser(ctx context.Context, id int32) (*dto.UserResponse, error) {
	var (
		logger = appContext.LoggerFromContext(ctx)
	)

	logger.Info("userUseCase.GetUser")
	user, err := u.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, appError.CodeErrGeneral.WithErrorDetail(err)
	}
	if user == nil {
		return nil, appError.CodeErrUserNotFound
	}
	userEntity := dto.UserResponse{}.FromUserEntity(user)
	return &userEntity, nil
}

func (u userUseCase) ListUsers(ctx context.Context, limit int, offset int) ([]dto.UserResponse, error) {
	var (
		logger = appContext.LoggerFromContext(ctx)
	)

	logger.Info("userUseCase.ListUsers")

	users, err := u.userRepository.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	var userList []dto.UserResponse
	for _, user := range users {
		userList = append(userList, dto.UserResponse{}.FromUserEntity(&user))
	}
	return userList, nil
}
