package interfaces

import entities "github.com/brianmorais/go-book-store-api/src/domain/entities"

type IUserRepository interface {
	GetUserByNameAndPassword(username, password string) (entities.User, error)
}
