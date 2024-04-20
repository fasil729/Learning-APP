package controllers

import (
	"Brilliant/application/dtos/exam_prep"
	"Brilliant/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExamPrepController struct {
	examPrepService *services.ExamPrepService
}

func NewExamPrepController(examPrepService *services.ExamPrepService) *ExamPrepController {
	return &ExamPrepController{
		examPrepService: examPrepService,
	}
}

// GenerateExamPrepMaterial godoc
// @Summary Generate exam preparation material
// @Description Generate exam preparation material based on the provided topics and prompt
// @Param generateDTO body dtos.GenerateExamPrepDTO true "Generate exam preparation material criteria"
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Tags exam_prep
// @Accept json
// @Produce json
// @Success 200 {string} string "Markdown formatted exam preparation material"
// @Failure 400 {object} dtos.ErrorResponse
// @Router /exam-prep/generate [post]
func (controller *ExamPrepController) GenerateExamPrepMaterial(ctx *gin.Context) {
	var generateDTO dtos.GenerateExamPrepDTO
	if err := ctx.ShouldBindJSON(&generateDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	material, err := controller.examPrepService.GenerateExamPrepMaterial(&generateDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusOK, material)
}