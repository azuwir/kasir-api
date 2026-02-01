package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (service *CategoryService) GetCategories() ([]models.Category, error) {
	return service.repo.GetCategories()
}

func (service *CategoryService) CreateCategory(category *models.Category) error {
	return service.repo.CreateCategory(category)
}

func (service *CategoryService) GetCategoryByID(id int) (*models.Category, error) {
	return service.repo.GetCategoryByID(id)
}

func (service *CategoryService) UpdateCategory(category *models.Category) error {
	return service.repo.UpdateCategory(category)
}

func (service *CategoryService) DeleteCategory(id int) error {
	return service.repo.DeleteCategory(id)
}
