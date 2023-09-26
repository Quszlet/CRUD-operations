package service

import (
	"github.com/Quszlet/CRUD-operations/internal/repository"
	"github.com/Quszlet/CRUD-operations/internal/models"
)

type User interface {
	Create(user models.User) (int, error)
	Update(userId int) error
	Get(userId int) (models.User, error)
	GetAll() ([]models.User, error)
	Delete(userId int) error
}

type Service struct {
	User
}

func NewService(r *repository.Repository) *Service {
	return &Service{User: NewUserService(r)}
}