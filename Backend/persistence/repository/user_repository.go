package repositories

import (
	contracts "Brilliant/application/contracts/persistence"
	"Brilliant/domain"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	GenericRepository[domain.User]
	database *gorm.DB
}

func NewUserRepository(getDb func() *gorm.DB) contracts.IUserRepository {
	return &UserRepository{
		database: getDb(),
	}
}

func (repository *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := repository.database.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *UserRepository) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := repository.database.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *UserRepository) FindByUsernameOrEmail(username, email string) (*domain.User, error) {
	user := &domain.User{}
	result := repository.database.Where("username = ? OR email = ?", username, email).First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repository *UserRepository) UpdateRefreshToken(userID uint, refreshToken string) error {
	return repository.database.Model(&domain.User{}).Where("id = ?", userID).Update("refresh_token", refreshToken).Error
}
