package service

import (
	"github.com/Quszlet/CRUD-operations/internal/models"
	"github.com/Quszlet/CRUD-operations/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(r *repository.Repository) *UserService {
	return &UserService{repo: r.User}
}

func (us *UserService) Create(user models.User) (int, error) {
	return us.repo.Create(user)
}

func (us *UserService) Update(userId int) error {
	return us.repo.Update(userId)
}

func (us *UserService) Get(userId int) (models.User, error) {
	return us.repo.Get(userId)
}

func (us *UserService) GetAll() ([]models.User, error) {
	return us.repo.GetAll()
}

func (us *UserService) Delete(userId int) error {
	return us.repo.Delete(userId)
}
