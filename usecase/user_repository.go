package usecase

import (
	"github.com/sozuuuuu/clean_architecture/domain"
)

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() ([]*domain.User, error)
}
