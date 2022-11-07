package repository

import (
	"demo-clean-arch/entity"
	"demo-clean-arch/entity/response"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user entity.User) error
	GetAll() ([]entity.User, error)
	GetUserbyID(id int64) (entity.User, error)
	DeleteUserbyID(id int64) error
	UpdateUser(user entity.User) (response.GetUserResponse, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) Create(user entity.User) error {
	if err := u.db.Debug().Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	if err := u.db.Debug().Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserRepository) GetUserbyID(id int64) (entity.User, error) {
	var user entity.User
	if err := u.db.Debug().Find(&user, id).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u UserRepository) DeleteUserbyID(id int64) error {
	var user entity.User
	if err := u.db.Debug().Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) UpdateUser(user entity.User) (response.GetUserResponse, error) {
	if err := u.db.Debug().Save(&user).Error; err != nil {
		return response.GetUserResponse{}, err
	}
	userResponse := response.GetUserResponse{}
	copier.Copy(&userResponse, &user)
	return userResponse, nil
}
