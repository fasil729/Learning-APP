package controllers

import (
	dtos "Brilliant/application/dtos/note"
	"Brilliant/application/services"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ChapterController handles HTTP requests related to chapters
type NoteController struct {
	noteService *services.NoteService
}

func NewNoteController(noteService *services.NoteService) *NoteController {
	return &NoteController{noteService: noteService}
}

// AddNoteWithImage handles HTTP request to add a note with text and image
// @Summary Add a note with text and image to a chapter
// @Description Add a note with text and image to the chapter specified by the chapter ID
// @Param chapterID path int true "Chapter ID"
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Tags notes
// @Accept mpfd
// @Produce json
// @Param text formData string true "Text content of the note"
// @Param image formData file false "Image file"
// @Success 200 {object} dtos.SuccessResponse "Note added successfully"
// @Failure 400 {object} dtos.ErrorResponse "Invalid chapter ID or missing
// @Router /chapters/{chapterID}/notes [post]
func (controller *NoteController) AddNoteWithImage(ctx *gin.Context) {
	chapterIDStr := ctx.Param("chapterID")
	chapterID, err := strconv.ParseUint(chapterIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	// Get the text content from the request
	text := ctx.PostForm("text")
	if text == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Text content is required"})
		return
	}

	// Get the image file from the request
	imageFile, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	// Read the image file and store it in a temporary location (for example)
	imageFileStream, err := imageFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer imageFileStream.Close()
	imageData, err := io.ReadAll(imageFileStream)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create the AddNoteDTO with the text and image data
	addNoteDTO := &dtos.AddNoteDTO{
		Text:      text,
		ImageData: imageData, // Image data as byte array
	}

	// Call the service method to add the note
	content, err := controller.noteService.AddNote(uint(chapterID), addNoteDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message Note added successfully": string(content)})
}

// GetNotesByChapterID handles HTTP request to get notes by chapter ID
// @Summary Get notes by chapter ID
// @Description Get all notes belonging to the chapter specified by the chapter ID
// @Param chapterID path int true "Chapter ID"
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Tags notes
// @Produce json
// @Success 200 {object} dtos.SuccessResponse "Notes retrieved successfully"
// @Failure 400 {object} dtos.ErrorResponse "Invalid chapter ID"
// @Router /chapters/{chapterID}/notes [get]
func (controller *NoteController) GetNoteByChapterID(ctx *gin.Context) {
	chapterIDStr := ctx.Param("chapterID")
	chapterID, err := strconv.ParseUint(chapterIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	note, err := controller.noteService.GetNoteByChapterID(uint(chapterID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"notes": string(note)})
}
