package usecase

import (
	"context"

	appContext "github.com/cde/go-example/src/context"
	appError "github.com/cde/go-example/src/error"
	userDTO "github.com/cde/go-example/src/modules/user/dto"
	"github.com/cde/go-example/src/modules/user/entity"
	userRepository "github.com/cde/go-example/src/modules/user/repository"
)

type userUseCase struct {
	userRepository userRepository.UserInterface
}

func NewUserUseCase(userRepository userRepository.UserInterface) UserUseCaseInterface {
	return &userUseCase{userRepository: userRepository}
}

func (u userUseCase) CreateUser(ctx context.Context, request *userDTO.UserRequest) (*userDTO.UserResponse, error) {
	var (
		logger = appContext.LoggerFromContext(ctx)
	)

	logger.Info("userUseCase.CreateUser")
	newRecord, err := u.userRepository.Create(
		ctx,
		&entity.User{
			Name:     request.Name,
			Email:    request.Email,
			Password: request.Password,
		},
	)
	if err != nil {
		return nil, err
	}

	return &userDTO.UserResponse{
		ID:    newRecord.ID,
		Name:  newRecord.Name,
		Email: newRecord.Email,
	}, err
}

func (u userUseCase) GetUser(ctx context.Context, id int32) (*userDTO.UserResponse, error) {
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
	userEntity := userDTO.UserResponse{}.FromUserEntity(user)
	return &userEntity, nil
}

func (u userUseCase) ListUsers(ctx context.Context, limit int, offset int) ([]userDTO.UserResponse, error) {
	var (
		logger = appContext.LoggerFromContext(ctx)
	)

	logger.Info("userUseCase.ListUsers")

	users, err := u.userRepository.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	var userList []userDTO.UserResponse
	for _, user := range users {
		userList = append(userList, userDTO.UserResponse{}.FromUserEntity(&user))
	}
	return userList, nil
}
