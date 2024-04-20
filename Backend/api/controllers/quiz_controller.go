package controllers

import (
	"Brilliant/application/dtos/quiz"
	"Brilliant/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type QuizController struct {
	quizService *services.QuizService
}

func NewQuizController(quizService *services.QuizService) *QuizController {
	return &QuizController{
		quizService: quizService,
	}
}

// GenerateQuiz godoc
// @Summary Generate a quiz
// @Description Generate a quiz based on the provided topics and prompt
// @Param generateDTO body dtos.GenerateQuizDTO true "Generate quiz criteria"
// @Tags quiz
// @Accept json
// @Produce json
// @Success 200 {array} dtos.Quiz
// @Failure 400 {object} dtos.ErrorResponse
// @Router /quiz/generate [post]
func (controller *QuizController) GenerateQuiz(ctx *gin.Context) {
	var generateDTO dtos.GenerateQuizDTO
	if err := ctx.ShouldBindJSON(&generateDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quizzes, err := controller.quizService.GenerateQuiz(&generateDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, quizzes)
}
