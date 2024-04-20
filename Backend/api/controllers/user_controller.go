package controllers

import (
	"Brilliant/application/dtos/user"
	"Brilliant/application/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with the provided information
// @Param user body dtos.RegisterDTO true "User registration information"
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} dtos.UserDTO
// @Failure 400 {object} dtos.ErrorResponse  
// @Router /users/register [post]
func (controller *UserController) RegisterUser(ctx *gin.Context) {
	var registerDTO dtos.RegisterDTO
	if err := ctx.ShouldBindJSON(&registerDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDTO, err := controller.userService.RegisterUser(&registerDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userDTO)
}

// RegisterAdmin godoc
// @Summary Register a new admin user
// @Description Register a new admin user with the provided information
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Param user body dtos.RegisterDTO true "Admin user registration information"
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} dtos.UserDTO
// @Failure 400 {object} dtos.ErrorResponse 
// @Router /users/admin/register [post]
func (controller *UserController) RegisterAdmin(ctx *gin.Context) {
	var registerDTO dtos.RegisterDTO
	if err := ctx.ShouldBindJSON(&registerDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDTO, err := controller.userService.RegisterAdmin(&registerDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userDTO)
}

// GetMe godoc
// @Summary Get user profile details
// @Description Retrieve the profile details of the currently authenticated user
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Tags users
// @Produce json
// @Success 200 {string} string "user profile details"
// @Router /users/me [get]
func (controller *UserController) GetMe(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "user profile details")
}

// SignIn godoc
// @Summary Sign in a user
// @Description Sign in a user with the provided credentials
// @Param user body dtos.SignInDTO true "User sign-in information"
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} dtos.UserDTO
// @Failure 400 {object} dtos.ErrorResponse  
// @Router /users/login [post]
func (controller *UserController) SignIn(ctx *gin.Context) {
	var signInDTO dtos.SignInDTO
	if err := ctx.ShouldBindJSON(&signInDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDTO, err := controller.userService.SignIn(&signInDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userDTO)
}

// UpdateUser godoc
// @Summary Update user information
// @Description Update the information of the currently authenticated user
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Param id path integer true "User ID"
// @Param user body dtos.UpdateUserDTO true "User update information"
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} dtos.UserDTO
// @Failure 400 {object} dtos.ErrorResponse 
// @Router /users/update/{id} [put]
func (controller *UserController) UpdateUser(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	var updateDTO dtos.UpdateUserDTO
	if err := ctx.ShouldBindJSON(&updateDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDTO, err := controller.userService.UpdateUser(userID.(uint), &updateDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	
	ctx.JSON(http.StatusOK, userDTO)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Param id path integer true "User ID"
// @Tags users
// @Produce json
// @Success 200 {object} dtos.SuccessResponse
// @Failure 400 {object} dtos.ErrorResponse  
// @Router /users/delete/{id} [delete]
func (controller *UserController) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	uintUserID, _ := strconv.ParseUint(userID, 10, 32)
	err := controller.userService.DeleteUser(uint(uintUserID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// Logout godoc
// @Summary Logout a user
// @Description Logout a user by invalidating the refresh token
// @Tags users
// @Produce json
// @Success 200 {object} dtos.SuccessResponse
// @Failure 400 {object} dtos.ErrorResponse  
// @Router /users/logout [post]
func (controller *UserController) Logout(ctx *gin.Context) {
	refreshToken, _ := ctx.Get("refresh_token")
	if refreshToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := controller.userService.Logout(refreshToken.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}


