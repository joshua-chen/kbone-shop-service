package services

import (
	"user/datamodels"
	"user/repositories"
)

type UserService interface {
	GetAll() []datamodels.User
	GetByID(id int64) (datamodels.User, bool)
	DeleteByID(id int64) bool
}

// NewUserService 返回默认的 user 服务层.
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositories.UserRepository
}

// GetAll 返回所有的 users.
func (s *userService) GetAll() []datamodels.User {
	return []datamodels.User{}
}

func (s *userService) GetByID(id int64) (datamodels.User, bool) {
	return datamodels.User{},true
}

func (s *userService) DeleteByID(id int64) bool {
	return true
}
