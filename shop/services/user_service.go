package services

import (
	"github.com/joshua-chen/go-commons/middleware/models"
	"github.com/joshua-chen/go-commons/utils/security/aes"
	"errors"
	"fmt"
	"shop/repositories"
	"time"

)

type UserService interface {
	GetAll() []models.User
	Registe(user *models.User) bool
	GetByID(id int64) (models.User, bool)
	DeleteByID(id int64) bool
}

// NewUserService 返回默认的 user 服务层.
func NewUserService() UserService {
	return &userService{
		repo: repositories.NewUserRepository(),
	}
}

type userService struct {
	repo repositories.UserRepository
}

// GetAll 返回所有的 users.
func (s *userService) GetAll() []models.User {
	return []models.User{}
}

func (s *userService) GetByID(id int64) (models.User, bool) {
	return models.User{},true
}

func (s *userService) DeleteByID(id int64) bool {
	return true
}

func (s *userService) Registe(user *models.User) bool {
	 

	user.Password = aes.AESEncrypt([]byte(user.Password))
	user.Enabled = 1
	user.CreateTime = time.Now()

	effect, err := s.repo.CreateUser(user)
	if effect <= 0 || err != nil {
		//ctx.Application().Logger().Errorf("用户[%s]注册失败。%s", user.Username, err.Error())
		//supports.Error(ctx, iris.StatusInternalServerError, supports.RegisteFailur, nil)
		panic(errors.New(fmt.Sprintf("用户[%s]注册失败。%s", user.Username)))
	}

	return true
}