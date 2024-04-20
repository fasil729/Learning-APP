package controllers

import (
	"Brilliant/application/dtos/experiment"
	"Brilliant/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

)

type ExperimentController struct {
	experimentService *services.ExperimentService
}

func NewExperimentController(experimentService *services.ExperimentService) *ExperimentController {
	return &ExperimentController{
		experimentService: experimentService,
	}
}

// GetExperimentsPerChapter godoc
// @Summary Get experiments for a chapter
// @Description Get a list of experiments for a specific chapter
// @Param chapterID path integer true "Chapter ID"
// @securityDefinitions.apiKey token
// @Param Authorization header string true "Authorization header" default(Bearer )
// @name Authorization
// @Security JWT
// @Tags experiments
// @Produce json
// @Success 200 {array} dtos.ExperimentDTO
// @Failure 400 
// @Router /experiments/chapter/{chapterID} [get]
func (controller *ExperimentController) GetExperimentsPerChapter(ctx *gin.Context) {
	chapterIDStr := ctx.Param("chapterID")
	chapterID, err := strconv.ParseUint(chapterIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapterID"})
		return
	}

	experiments, err := controller.experimentService.GetExperimentsPerChapter(uint(chapterID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, experiments)
}

// GetExperimentsPerSubject godoc
// @Summary Get experiments for a subject
// @Description Get a list of experiments for a specific subject
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Param subjectID path integer true "Subject ID"
// @Tags experiments
// @Produce json
// @Success 200 {array} dtos.ExperimentDTO
// @Failure 400 {object} dtos.ErrorResponse 
// @Router /experiments/subject/{subjectID} [get]
func (controller *ExperimentController) GetExperimentsPerSubject(ctx *gin.Context) {
	subjectIDStr := ctx.Param("subjectID")
	subjectID, err := strconv.ParseUint(subjectIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subjectID"})
		return
	}

	experiments, err := controller.experimentService.GetExperimentsPerSubject(uint(subjectID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, experiments)
}


// GetExperimentContent godoc
// @Summary Get content for an experiment
// @Description Get the content for a specific experiment
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Param experimentID path integer true "Experiment ID"
// @Tags experiments
// @Produce json
// @Success 200 {object} dtos.SuccessResponse
// @Failure 400 {object} dtos.ErrorResponse 
// @Router /experiments/content/{experimentID} [get]
func (controller *ExperimentController) GetExperimentContent(ctx *gin.Context) {
	experimentIDStr := ctx.Param("experimentID")
	experimentID, err := strconv.ParseUint(experimentIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experimentID"})
		return
	}

	content, err := controller.experimentService.GetExperimentContent(uint(experimentID), "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"content": string(content)})
}

// GenerateExperimentsPerChapter godoc
// @Summary Generate experiments for a chapter
// @Description Generate experiments for a specific chapter based on provided criteria
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Param generateDTO body dtos.GenerateExperimentDTO true "Generate experiment criteria"
// @Tags experiments
// @Accept json
// @Produce json
// @Success 200 {array} dtos.ExperimentDTO
// @Failure 400 {object} dtos.ErrorResponse  
// @Router /experiments/generate [post]
func (controller *ExperimentController) GenerateExperimentsPerChapter(ctx *gin.Context) {
	var generateDTO dtos.GenerateExperimentDTO
	if err := ctx.ShouldBindJSON(&generateDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	experiments, err := controller.experimentService.GenerateExperimentsPerChapter(&generateDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, experiments)
}

// AddExperiment godoc
// @Summary Add a new experiment
// @Description Add a new experiment to the database
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Param addDTO body dtos.ExperimentDTO true "New experiment data"
// @Tags experiments
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ExperimentDTO
// @Failure 400 {object} dtos.ErrorResponse  
// @Router /experiments/add [post]
func (controller *ExperimentController) AddExperiment(ctx *gin.Context) {
	var addDTO dtos.ExperimentDTO
	if err := ctx.ShouldBindJSON(&addDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	experiment, err := controller.experimentService.CreateExperiment(&addDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, experiment)
}

// UpdateExperiment godoc
// @Summary Update an existing experiment
// @Description Update an existing experiment in the database
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Param experimentID path integer true "Experiment ID"
// @Param updateDTO body dtos.ExperimentDTO true "Updated experiment data"
// @Tags experiments
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ExperimentDTO
// @Failure 400 {object} dtos.ErrorResponse 
// @Router /experiments/update/{experimentID} [put]
func (controller *ExperimentController) UpdateExperiment(ctx *gin.Context) {
	experimentIDStr := ctx.Param("experimentID")
	experimentID, err := strconv.ParseUint(experimentIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experimentID"})
		return
	}

	var updateDTO dtos.ExperimentDTO
	if err := ctx.ShouldBindJSON(&updateDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	experiment, err := controller.experimentService.UpdateExperiment(uint(experimentID), &updateDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, experiment)
}

// DeleteExperiment godoc
// @Summary Delete an experiment
// @Description Delete an experiment from the database
// @Param Authorization header string true "Authorization header" default(Bearer )
// @Param experimentID path integer true "Experiment ID"
// @Tags experiments
// @Produce json
// @Success 200 {object} dtos.SuccessResponse
// @Failure 400 {object} dtos.ErrorResponse 
// @Router /experiments/delete/{experimentID} [delete]
func (controller *ExperimentController) DeleteExperiment(ctx *gin.Context) {
	experimentIDStr := ctx.Param("experimentID")
	experimentID, err := strconv.ParseUint(experimentIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experimentID"})
		return
	}

	if err := controller.experimentService.DeleteExperiment(uint(experimentID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Experiment deleted successfully"})
}
