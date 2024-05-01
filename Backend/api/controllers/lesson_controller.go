package controllers

import (
	dtos "Brilliant/application/dtos/lesson"
	"Brilliant/application/services"
	"fmt"
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
// @Description Get the content of a lesson by Name
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Tags lessons
// @Accept  json
// @Produce  json
// @Param lessonName path string true "Lesson Name"
// @Success 200 {object} dtos.SuccessResponse
// @Failure 400 {object} dtos.ErrorResponse "Bad request"
// @Router /lessons/{lessonName}/content [get]
func (controller *LessonController) GetLessonContent(ctx *gin.Context) {
	lessonName := ctx.Param("lessonName")
	fmt.Println("At the controller", lessonName)
	// lessonID, err := strconv.ParseUint(lessonIDstr, 10, 64)

	lessonContent, err := controller.lessonService.GetLessonContent(lessonName)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Could not retrieve lesson content. Please try again later."})
		return
	}

	// If content is text, set the appropriate headers
	ctx.Header("Content-Type", "text/plain")
	ctx.String(200, string(lessonContent))
}

// @Summary Create a new lesson
// @Description Create a new lesson with the provided information
// @Param Authorization header string true "Authorization header" default(Bearer )
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
// @Param Authorization header string true "Authorization header" default(Bearer )
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
// @Param Authorization header string true "Authorization header" default(Bearer )
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
