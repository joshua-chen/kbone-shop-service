package services

import (
	"github.com/joshua-chen/go-commons/exception"
	"github.com/joshua-chen/go-commons/middleware/jwt"
	_"github.com/joshua-chen/go-commons/mvc/context/request"
	"github.com/joshua-chen/go-commons/middleware/models"
	_"github.com/joshua-chen/go-commons/utils/security/aes"
	_"errors"
	_"fmt"
	_"shop/datamodels"
	"shop/repositories"
	_"time"

)

type SessionService interface {

	NewToken(userId int64,username string) string
}

// NewSessionService 返回默认的 Session 服务层.
func NewSessionService() SessionService {
	return &sessionService{
		repo: repositories.NewUserRepository(),
	}
}

type sessionService struct {
	repo repositories.UserRepository
}


func (s *sessionService) NewToken(userId int64,username string) string {

	user:= models.User{Id:userId,Username:username}
	token, err:= jwt.NewToken(&user)
	if err != nil{
		exception.Fatal(err)
	}

	return token
}