package repository

import (
	"github.com/Quszlet/CRUD-operations/internal/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(user models.User) (int, error)
	Update(userId int) error
	Get(userId int) (models.User, error)
	GetAll() ([]models.User, error)
	Delete(userId int) error
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{User: NewUserPostgres(db)}
}
