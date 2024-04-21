package contracts

import "Brilliant/domain"

type IUserRepository interface {
	IGenericRepository[domain.User]
	GetByEmail(email string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	FindByUsernameOrEmail(username string, email string) (*domain.User, error)
	UpdateRefreshToken(userID uint, refreshToken string) error
}
