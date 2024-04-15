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

func (controller *UserController) GetMe(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "user profile details")
}

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


