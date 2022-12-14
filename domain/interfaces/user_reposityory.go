package interfaces

import "github.com/matheusrbarbosa/squilo/domain/models"

type UserRepository interface {
	Create(user models.User) models.User
	GetByEmail(email string) (models.User, error)
	GetById(id int) (models.User, error)
}
