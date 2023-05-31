package user

import (
	"github.com/OneeyK/KPI_GoLang/pkg/models"
)

type userRepository interface {
	FindByID(userID uint64) (*models.User, error)
	Remove(userID uint64) error
	Update(userID uint64, updatedUser models.User) (*models.User, error)
	GetAll() ([]models.User, error)
}

type UserService struct {
	userRepository userRepository
}

func NewUserService(userRepository userRepository) User {
	return &UserService{userRepository: userRepository}
}

type User interface {
	GetByID(id uint64) (*models.User, error)
	DeleteByID(id uint64) error
	Update(userID uint64, updatedUser models.User) (*models.User, error)
	GetAll() ([]models.User, error)
}

func (s *UserService) GetByID(id uint64) (*models.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *UserService) DeleteByID(id uint64) error {
	return s.userRepository.Remove(id)
}

func (s *UserService) Update(userID uint64, updatedUser models.User) (*models.User, error) {
	return s.userRepository.Update(userID, updatedUser)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.userRepository.GetAll()
}
