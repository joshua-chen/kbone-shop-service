/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:11:13
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 00:28:56
 */
package repositories

import (
	"github.com/joshua-chen/go-commons/datasource"
	"github.com/joshua-chen/go-commons/middleware/models"
	"github.com/joshua-chen/go-commons/mvc/context/request"
	_ "errors"
	"sync"

)

// Query代表一种“访客”和它的查询动作。
type UserQuery func(models.User) bool

// UserRepository会处理一些关于user实例的基本的操作 。
// 这是一个以测试为目的的接口，即是一个内存中的user库
// 或是一个连接到数据库的实例。
type UserRepository interface {
	CreateUser(user ...*models.User) (int64, error)

	GetUserByUsername(user *models.User) (bool, error)

	GetUsersByUids(uids []int64, page *request.Pagination) ([]*models.User, int64, error)

	UpdateUserById(user *models.User) (int64, error)

	DeleteByUsers(uids []int64) (effect int64, err error)

	GetPaginationUsers(page *request.Pagination) ([]*models.User, int64, error)
}

// NewUserRepository返回一个新的基于内存的user库。
// 库的类型在我们的例子中是唯一的。
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// userMemoryRepository就是一个"UserRepository"
// 它负责存储于内存中的实例数据(map)
type userRepository struct {
	mu sync.RWMutex
}

func (repo *userRepository) CreateUser(user ...*models.User) (int64, error) {
	e := datasource.MasterEngine()
	return e.Insert(user)
}

func (repo *userRepository) GetUserByUsername(user *models.User) (bool, error) {
	e := datasource.MasterEngine()
	return e.Get(user)
}

func (repo *userRepository) GetUsersByUids(uids []int64, page *request.Pagination) ([]*models.User, int64, error) {
	e := datasource.MasterEngine()
	users := make([]*models.User, 0)

	s := e.In("id", uids).Limit(page.Limit, page.Offset)
	if page.SortName != "" {
		switch page.SortOrder {
		case "asc":
			s.Asc(page.SortName)
		case "desc":
			s.Desc(page.SortName)
		}
	}
	count, err := s.FindAndCount(&users)
	return users, count, err
}

func (repo *userRepository) UpdateUserById(user *models.User) (int64, error) {
	e := datasource.MasterEngine()
	return e.ID(user.Id).Update(user)
}

func (repo *userRepository) DeleteByUsers(uids []int64) (effect int64, err error) {
	e := datasource.MasterEngine()

	u := new(models.User)
	for _, v := range uids {
		i, err1 := e.ID(v).Delete(u)
		effect += i
		err = err1
	}
	return
}

func (repo *userRepository) GetPaginationUsers(page *request.Pagination) ([]*models.User, int64, error) {
	e := datasource.MasterEngine()
	userList := make([]*models.User, 0)

	s := e.Limit(page.Limit, page.Offset)
	if page.SortName != "" {
		switch page.SortOrder {
		case "asc":
			s.Asc(page.SortName)
		case "desc":
			s.Desc(page.SortName)
		}
	}
	count, err := s.FindAndCount(&userList)

	return userList, count, err
}
