package repositories

import (
    "gorm.io/gorm"
    "hackathon.com/leariningApp/domain"
)

type UserRepository struct {
    *GenericRepository[domain.User]
}

func NewUserRepository(getDb func() *gorm.DB) *UserRepository {
    return &UserRepository{
        GenericRepository: NewGenericRepository[domain.User](getDb),
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
