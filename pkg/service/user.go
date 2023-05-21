package service

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetByID(id uint64) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) DeleteByID(id uint64) error {
	return s.repo.Remove(id)
}

func (s *UserService) Update(userID uint64, updatedUser models.User) (*models.User, error) {
	updatedUser.Password = generatePasswordHash(updatedUser.Password)
	return s.repo.Update(userID, updatedUser)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}
