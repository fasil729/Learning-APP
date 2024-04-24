package controllers

import (
	"Brilliant/application/dtos/lesson"
	"Brilliant/application/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LessonController struct {
	lessonService *services.LessonService
}

func NewLessonController(lessonService *services.LessonService) *LessonController {
	return &LessonController{lessonService: lessonService}
}

// @Summary Get lesson content
// @Description Get the content of a lesson by ID
// @Tags lessons
// @Accept  json
// @Produce  json
// @Param lessonID path int true "Lesson ID"
// @Success 200 {object} dtos.SuccessResponse
// @Failure 400 {object} dtos.ErrorResponse "Bad request"
// @Router /lessons/{lessonID}/content [get]
func (controller *LessonController) GetLessonContent(ctx *gin.Context) {
	lessonIDstr := ctx.Param("lessonID")
	lessonID, err := strconv.ParseUint(lessonIDstr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	lesson, err := controller.lessonService.GetLessonContent(uint(lessonID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, lesson)
}

// @Summary Create a new lesson
// @Description Create a new lesson with the provided information
// @Tags lessons
// @Accept  json
// @Produce  json
// @Param lesson body dtos.LessonDTO true "Lesson information"
// @Success 201 {object} dtos.SuccessResponse
// @Failure 400 {object} dtos.ErrorResponse "Bad request"
// @Router /lessons [post]
func (controller *LessonController) AddLesson(ctx *gin.Context) {
	lessonDTO := dtos.LessonDTO{}
	if err := ctx.ShouldBindJSON(lessonDTO); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	lesson, err := controller.lessonService.CreateLesson(&lessonDTO)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, lesson)
}

// @Summary Update a lesson
// @Description Update the information of a lesson by ID
// @Tags lessons
// @Accept  json
// @Produce  json
// @Param lessonID path int true "Lesson ID"
// @Param lesson body dtos.LessonDTO true "Lesson information"
// @Success 200 {object} dtos.SuccessResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Router /lessons/{lessonID} [put]
func (controller *LessonController) UpdateLesson(ctx *gin.Context) {
	lessonIDstr := ctx.Param("lessonID")
	lessonID, err := strconv.ParseUint(lessonIDstr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	lessonDTO := dtos.LessonDTO{}
	if err := ctx.ShouldBindJSON(lessonDTO); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	lesson, err := controller.lessonService.UpdateLesson(uint(lessonID), &lessonDTO)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, lesson)
}

// @Summary Delete a lesson
// @Description Delete a lesson by ID
// @Tags lessons
// @Accept  json
// @Produce  json
// @Param lessonID path int true "Lesson ID"
// @Success 204 {object} dtos.SuccessResponse "Successfully deleted the lesson"
// @Failure 400 {object} dtos.ErrorResponse "Bad request"
// @Router /lessons/{lessonID} [delete]
func (controller *LessonController) DeleteLesson(ctx *gin.Context) {
	lessonIDstr := ctx.Param("lessonID")
	lessonID, err := strconv.ParseUint(lessonIDstr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = controller.lessonService.DeleteLesson(uint(lessonID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Lesson deleted successfully"})
}
