package controllers

import (
	dtos "Brilliant/application/dtos/subject"
	"Brilliant/application/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubjectController struct {
	subjectService *services.SubjectService
}

func NewSubjectController(subjectService *services.SubjectService) *SubjectController {
	return &SubjectController{subjectService: subjectService}
}

// CreateSubject godoc
// @Summary Create a new subject
// @Description Create a new subject with the provided information
// @Param subject body dtos.CreateSubjectDTO true "Subject creation information"
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Tags subjects
// @Accept json
// @Produce json
// @Success 200 {object} dtos.SuccessResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Router /subjects/create [post]
func (controller *SubjectController) CreateSubject(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	var subjectDTO *dtos.CreateSubjectDTO
	if err := ctx.ShouldBindJSON(&subjectDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject, err := controller.subjectService.CreateSubject(userID.(uint), subjectDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subject)
}

// SearchSubjectsByName godoc
// @Summary Search subjects by name
// @Description Search for subjects by name
// @Param query query string true "Search query"
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Tags subjects
// @Produce json
// @Success 200 {array} dtos.SuccessResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /subjects/search [get]
func (controller *SubjectController) SearchSubjectsByName(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	query := ctx.Query("query")
	subjects, err := controller.subjectService.SearchSubjectsByName(userID.(uint), query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subjects)
}
