/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-18 14:54:08
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-18 15:00:43
 */ 
package services

import (
	"shop/datamodels"
	"shop/repositories"

)

type ProductService interface {
	GetAll() []datamodels.Product
	GetByID(id int64) (datamodels.Product, bool)
	DeleteByID(id int64) bool
}

// NewProductService 返回默认的 product 服务层.
func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

type productService struct {
	repo repositories.ProductRepository
}

// GetAll 返回所有的 products.
func (s *productService) GetAll() []datamodels.Product {
	return s.repo.SelectMany(func(_ datamodels.Product) bool {
		return true
	}, -1)
}

func (s *productService) GetByID(id int64) (datamodels.Product, bool) {
	return s.repo.Select(func(m datamodels.Product) bool {
		return m.ID == id
	})
}

func (s *productService) DeleteByID(id int64) bool {
	return true
}
