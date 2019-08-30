package services

import (
	"github.com/iamMarkchu/goldfish/controllers/requests"
	"github.com/iamMarkchu/goldfish/models"
)

type CategoryService struct {
	*models.Category
}

func NewCategoryService(category *models.Category) *CategoryService {
	return &CategoryService{Category: category}
}

func (s *CategoryService) Create(scr requests.StoreCategoryRequest, userId int) (*models.Category, error) {
	var err error
	s.Category.Name = scr.Name
	s.Category.AliasName = scr.AliasName
	s.Category.DisplayOrder = scr.DisplayOrder
	s.Category.ParentCat = &models.Category{Id:scr.ParentId}
	s.Category.Image = scr.Image
	s.Category.Status = models.StatusNormal
	s.Category.User = &models.User{Id: userId}

	if err = s.Category.Create(); err != nil {
		return nil, err
	}
	return s.Category, nil
}

func (s *CategoryService) GetCategoryById(i int) (*models.Category, error) {
	var err error
	s.Category.Id = i
	if err = s.Category.GetCategoryById(); err != nil {
		return nil, err
	}
	return s.Category, nil
}

func (s *CategoryService) Update(scr requests.StoreCategoryRequest, userId int, id int) (*models.Category, error) {
	var err error
	s.Category.Id = id
	s.Category.GetCategoryById()

	if err = s.Category.Update(); err != nil {
		return nil, err
	}
	return s.Category, nil
}
