package controllers

import (
	"github.com/gin-gonic/gin"
	"hackathon.com/leariningApp/application/services"
	"hackathon.com/leariningApp/domain"
	"net/http"
	"strconv"
)

type SubjectController struct {
	subjectService services.SubjectService
}

func NewSubjectController(subjectService services.SubjectService) *SubjectController {
	return &SubjectController{subjectService: subjectService}
}

func (controller *SubjectController) CreateSubject(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var subject *domain.Subject
	if err := ctx.ShouldBindJSON(&subject); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject, err = controller.subjectService.CreateSubject(subject.SubjectName, uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subject)
}

func (controller *SubjectController) SearchSubjectsByName(ctx *gin.Context) {
	query := ctx.Query("query")
	subjects, err := controller.subjectService.SearchSubjectsByName(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subjects)
}

func (controller *SubjectController) GenerateRoadMap(ctx *gin.Context) {
	var subject domain.Subject
	if err := ctx.ShouldBindJSON(&subject); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chapters, lessons, err := controller.subjectService.GenerateRoadMap(&subject)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"chapters": chapters, "lessons": lessons})
}
