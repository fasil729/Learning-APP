package contracts

import (
	dtos "Brilliant/application/dtos/experiment"
	"Brilliant/domain"
)

// IExperimentHandler defines the contract for handling experiments in Gemini.
type IExperimentHandler interface {
	GetExperimentsForChapter(generateDTO *dtos.GenerateExperimentDTO, chapterName string) ([]domain.Experiment, error)
	GetExperimentContent(experimentName string, promtMessage string) ([]byte, error)
}