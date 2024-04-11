package repositories

import "hackathon.com/leariningApp/domain"

type IUserRepository interface {
	IGenericRepository[domain.User]
	GetByEmail(email string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
}
