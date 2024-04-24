package controllers

import (
	"Brilliant/application/dtos/chapter"
	"Brilliant/application/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ChapterController struct {
	chapterService *services.ChapterService
}

func NewChapterController(chapterService *services.ChapterService) *ChapterController {
	return &ChapterController{chapterService: chapterService}

}

// @Summary Create a new chapter
// @Description Create a new chapter with the provided information
// @Tags chapters
// @Accept  json
// @Produce  json
// @Param chapter body dtos.ChapterDTO true "Chapter information"
// @Success 201 {object} dtos.SuccessResponse "Successfully created the chapter"
// @Failure 400 {object} dtos.ErrorResponse "Bad request"
// @Router /chapters [post]
func (controller *ChapterController) CreateChapter(ctx *gin.Context) {
	var chapterDTO dtos.ChapterDTO
	if err := ctx.ShouldBindJSON(&chapterDTO); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	chapter, err := controller.chapterService.CreateChapter(&chapterDTO)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, chapter)
}

// @Summary Update a chapter
// @Description Update the information of a chapter by ID
// @Tags chapters
// @Accept  json
// @Produce  json
// @Param chapterID path int true "Chapter ID"
// @Param chapter body dtos.ChapterDTO true "Chapter information"
// @Success 200 {object} dtos.SuccessResponse "Successfully updated the chapter"
// @Failure 400 {object} dtos.ErrorResponse "Bad request"
// @Router /chapters/{chapterID} [put]
func (controller *ChapterController) UpdateChapter(ctx *gin.Context) {
	chapterIdStr := ctx.Param("chapterId")
	chapterId, err := strconv.ParseUint(chapterIdStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var chapterDTO dtos.ChapterDTO
	if err := ctx.ShouldBindJSON(&chapterDTO); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	chapter, err := controller.chapterService.UpdateChapter(uint(chapterId), &chapterDTO)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, chapter)
}

// @Summary Delete a chapter
// @Description Delete a chapter by ID
// @Tags chapters
// @Accept  json
// @Produce  json
// @Param chapterID path int true "Chapter ID"
// @Success 200 {object} dtos.SuccessResponse "Successfully deleted the chapter"
// @Failure 400 {object} dtos.ErrorResponse "Bad request"
// @Router /chapters/{chapterID} [delete]
func (controller *ChapterController) DeleteChapter(ctx *gin.Context) {
	chapterIdStr := ctx.Param("chapterId")
	chapterId, err := strconv.ParseUint(chapterIdStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = controller.chapterService.DeleteChapter(uint(chapterId))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Chapter deleted successfully"})
}
