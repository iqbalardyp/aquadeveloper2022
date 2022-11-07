package usecase

import (
	"demo-clean-arch/entity"
	"demo-clean-arch/entity/response"
	"demo-clean-arch/repository"

	"github.com/jinzhu/copier"
)

type IUserUseCase interface {
	CreateUser(user response.CreateUserRequest) error
	GetListUser() ([]response.GetUserResponse, error)
	GetByID(id int64) (response.GetUserResponse, error)
	DeleteByID(id int64) error
	UpdateByID(req response.GetUserResponse, id int64) (response.GetUserResponse, error)
}

type UserUseCase struct {
	userRepository repository.IUserRepository
}

func NewUserUseCase(userRepository repository.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepository: userRepository}
}

func (u UserUseCase) CreateUser(req response.CreateUserRequest) error {
	users := entity.User{}
	copier.Copy(&users, &req)

	err := u.userRepository.Create(users)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUseCase) GetList() ([]response.GetUserResponse, error) {
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	userResponse := []response.GetUserResponse{}
	copier.Copy(&userResponse, &users)
	return userResponse, nil
}

func (u UserUseCase) GetByID(id int64) (response.GetUserResponse, error) {
	user, err := u.userRepository.GetUserbyID(id)
	if err != nil {
		return response.GetUserResponse{}, err
	}
	userResponse := response.GetUserResponse{}
	copier.Copy(&userResponse, &user)
	return userResponse, nil
}

func (u UserUseCase) DeleteByID(id int64) error {
	if err := u.userRepository.DeleteUserbyID(id); err != nil {
		return err
	}
	return nil
}

func (u UserUseCase) UpdateByID(req response.GetUserResponse, id int64) (response.GetUserResponse, error) {
	user, err := u.userRepository.GetUserbyID(id)
	if err != nil {
		return response.GetUserResponse{}, err
	}
	copier.Copy(&user, &req)
	userResponse, err := u.userRepository.UpdateUser(user)
	if err != nil {
		return response.GetUserResponse{}, err
	}
	return userResponse, err
}
