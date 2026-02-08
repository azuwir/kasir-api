package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (service *ProductService) GetProducts(name string) ([]models.Product, error) {
	return service.repo.GetProducts(name)
}

func (service *ProductService) CreateProduct(product *models.Product) error {
	return service.repo.CreateProduct(product)
}

func (service *ProductService) GetProductByID(id int) (*models.Product, error) {
	return service.repo.GetProductByID(id)
}

func (service *ProductService) UpdateProduct(product *models.Product) error {
	return service.repo.UpdateProduct(product)
}

func (service *ProductService) DeleteProduct(id int) error {
	return service.repo.DeleteProduct(id)
}
