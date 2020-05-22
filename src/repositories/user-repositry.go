/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:11:13
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-18 15:33:22
 */
package repositories

import (
	_ "errors"
	"sync"
	"shop/datamodels"

)

// Query代表一种“访客”和它的查询动作。
type UserQuery func(datamodels.User) bool

// UserRepository会处理一些关于user实例的基本的操作 。
// 这是一个以测试为目的的接口，即是一个内存中的user库
// 或是一个连接到数据库的实例。
type UserRepository interface {
}

// NewUserRepository返回一个新的基于内存的user库。
// 库的类型在我们的例子中是唯一的。
func NewUserRepository(source map[int64]datamodels.User) UserRepository {
	return &userMemoryRepository{source: source}
}

// userMemoryRepository就是一个"UserRepository"
// 它负责存储于内存中的实例数据(map)
type userMemoryRepository struct {
	source map[int64]datamodels.User
	mu     sync.RWMutex
}
