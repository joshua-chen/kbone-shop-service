/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-18 14:54:08
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 12:14:20
 */ 
package services

import (
	"shop/datamodels"
	"shop/repositories"
	_ "commons/mvc/models"
	"commons/mvc/context/request"

)

type ProductService interface {
	GetAll(page *request.Pagination) ([]*datamodels.Product ,int64)
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
func (s *productService) GetAll( page *request.Pagination) ([]*datamodels.Product ,int64){
	
	return s.repo.SelectMany(page)
}

func (s *productService) GetByID(id int64) (datamodels.Product, bool) {
	return datamodels.Product{},true
}

func (s *productService) DeleteByID(id int64) bool {
	return true
}
