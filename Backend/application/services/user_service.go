package services

import (
	contracts "Brilliant/application/contracts/persistence"
	dtos "Brilliant/application/dtos/user"
	"Brilliant/application/util"
	"Brilliant/domain"
	"errors"
)

type UserService struct {
	userRepository  contracts.IUserRepository
	refreshTokenExp int
	accessTokenExp  int
	secret          string
}

func (service *UserService) RegisterUser(dto *dtos.RegisterDTO) (*dtos.UserDTO, error) {
	// Check if the role is valid
	if dto.Role != "user" && dto.Role != "student" {
		return nil, errors.New("invalid role")
	}

	// Check if there is already a user with the given username or email
	existingUser, err := service.userRepository.FindByUsernameOrEmail(dto.Username, dto.Email)
	if existingUser != nil {
		return nil, errors.New("username or email already exists")
	}

	// Hash the password
	hashedPassword, err := util.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	// Create the user entity
	user := &domain.User{
		Username: dto.Username,
		Password: hashedPassword,
		Email:    dto.Email,
		Role:     domain.UserRole(dto.Role),
	}

	// Register user logic here
	createdUser, err := service.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	// Create access and refresh tokens
	accessToken, err := util.CreateAccessToken(createdUser.ID, dto.Username, dto.Email, dto.Role, service.secret, service.accessTokenExp)
	if err != nil {
		return nil, err
	}
	refreshToken, err := util.CreateRefreshToken(createdUser.ID, dto.Username, dto.Email, dto.Role, service.secret, service.refreshTokenExp)
	if err != nil {
		return nil, err
	}

	// Store the refresh token in the database for later use
	err = service.userRepository.UpdateRefreshToken(createdUser.ID, refreshToken)
	if err != nil {
		return nil, err
	}

	userDTO := &dtos.UserDTO{
		ID:           createdUser.ID,
		Username:     createdUser.Username,
		Email:        createdUser.Email,
		Role:         createdUser.Role,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return userDTO, nil
}

func (service *UserService) RegisterAdmin(dto *dtos.RegisterDTO) (*dtos.UserDTO, error) {
	// Check if the role is valid
	if dto.Role != "admin" {
		return nil, errors.New("invalid role")
	}

	// Check if there is already a user with the given username or email
	existingUser, err := service.userRepository.FindByUsernameOrEmail(dto.Username, dto.Email)
	if existingUser != nil {
		return nil, errors.New("username or email already exists")
	}

	// Hash the password
	hashedPassword, err := util.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	// Create the user entity
	user := &domain.User{
		Username: dto.Username,
		Password: hashedPassword,
		Email:    dto.Email,
		Role:     domain.UserRole(dto.Role),
	}

	// Register user logic here
	createdUser, err := service.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	// Create access and refresh tokens
	accessToken, err := util.CreateAccessToken(createdUser.ID, dto.Username, dto.Email, dto.Role, service.secret, service.accessTokenExp)
	if err != nil {
		return nil, err
	}
	refreshToken, err := util.CreateRefreshToken(createdUser.ID, dto.Username, dto.Email, dto.Role, service.secret, service.refreshTokenExp)
	if err != nil {
		return nil, err
	}

	// Store the refresh token in the database for later use
	err = service.userRepository.UpdateRefreshToken(createdUser.ID, refreshToken)
	if err != nil {
		return nil, err
	}

	userDTO := &dtos.UserDTO{
		ID:           createdUser.ID,
		Username:     createdUser.Username,
		Email:        createdUser.Email,
		Role:         createdUser.Role,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return userDTO, nil
}

func (service *UserService) SignIn(dto *dtos.SignInDTO) (*dtos.UserDTO, error) {
	// Find the user by username
	user, err := service.userRepository.GetByUsername(dto.Username)
	if err != nil {
		return nil, err
	}

	// Check if the password matches
	if !util.CheckPassword(user.Password, dto.Password) {
		return nil, errors.New("invalid username or password")
	}

	// Create access and refresh tokens
	accessToken, err := util.CreateAccessToken(user.ID, user.Username, user.Email, string(user.Role), service.secret, service.accessTokenExp)
	if err != nil {
		return nil, err
	}
	refreshToken, err := util.CreateRefreshToken(user.ID, user.Username, user.Email, string(user.Role), service.secret, service.refreshTokenExp)
	if err != nil {
		return nil, err
	}

	// Update the refresh token in the database
	err = service.userRepository.UpdateRefreshToken(user.ID, refreshToken)
	if err != nil {
		return nil, err
	}

	// Create and return the user DTO with tokens
	userDTO := &dtos.UserDTO{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Role:         user.Role,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return userDTO, nil
}

func (service *UserService) UpdateUser(userID uint, dto *dtos.UpdateUserDTO) (*dtos.UserDTO, error) {
	// Find the user by ID
	user, err := service.userRepository.GetById(userID)
	if err != nil {
		return nil, err
	}

	// Update user fields
	user.Username = dto.Username
	user.Email = dto.Email
	// Update other fields as needed

	// Save the updated user
	updatedUser, err := service.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	// Create and return the updated user DTO
	updatedUserDTO := &dtos.UserDTO{
		ID:       updatedUser.ID,
		Username: updatedUser.Username,
		Email:    updatedUser.Email,
		Role:     updatedUser.Role,
		// Add other fields as needed
	}

	return updatedUserDTO, nil
}

func (service *UserService) DeleteUser(userID uint) error {
	user, err := service.userRepository.GetById(userID)
	if err != nil {
		return err
	}

	// Delete the user
	_, err = service.userRepository.Delete(user)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserService) Logout(refreshToken string) error {
    // Verify the refresh token
    userID, err := util.ExtractUserID(refreshToken, service.secret)
    if err != nil {
        return err
    }

    // Update the refresh token in the database to invalidate it
    err = service.userRepository.UpdateRefreshToken(userID, "")
    if err != nil {
        return err
    }

    return nil
}


// NewUserFeatures creates a new instance of UserFeatures.
func NewUserService(userRepository contracts.IUserRepository, refreshTokenExp, accessTokenExp int, secret string) *UserService {
	return &UserService{
		userRepository:  userRepository,
		refreshTokenExp: refreshTokenExp,
		accessTokenExp:  accessTokenExp,
		secret:          secret,
	}
}
