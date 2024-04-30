package services

import (
	contract "Brilliant/application/contracts/gemini"
	contracts "Brilliant/application/contracts/persistence"
	dtos "Brilliant/application/dtos/experiment"
	"Brilliant/domain"
)

type ExperimentService struct {
	experimentRepository   contracts.IExperimentRepository
	expermentgeminiHandler contract.IExperimentHandler
}

func NewExperimentService(experimentRepository contracts.IExperimentRepository, expermentgeminiHandler contract.IExperimentHandler) *ExperimentService {
	return &ExperimentService{
		experimentRepository:   experimentRepository,
		expermentgeminiHandler: expermentgeminiHandler,
	}
}

func (service *ExperimentService) GetExperimentsPerChapter(chapterID uint) ([]*domain.Experiment, error) {
	experiments, err := service.experimentRepository.GetByChapterID(chapterID)
	if err != nil {
		return nil, err
	}
	return experiments, nil
}

func (service *ExperimentService) GetExperimentsPerSubject(subjectID uint) ([]*domain.Experiment, error) {
	experiments, err := service.experimentRepository.GetBySubjectID(subjectID)
	if err != nil {
		return nil, err
	}
	return experiments, nil
}

func (service *ExperimentService) GetExperimentContent(experimentID uint, promptMessage string) ([]byte, error) {
	experiment, err := service.experimentRepository.GetById(experimentID)
	if err != nil {
		return nil, err
	}

	// if experiment.ContentLink == "" {
	// 	// Gemini API call will be done here

	// 	// For now, return a dummy experiment
	// 	experiment.ContentLink = "data/experiment/generated_experiment.json"

	// 	// Update the experiment with the content link
	// 	if _, err := service.experimentRepository.Update(experiment); err != nil {
	// 		return nil, err
	// 	}
	// }

	// // Read the content of the JSON file
	// content, err := os.ReadFile(experiment.ContentLink)
	// if err != nil {
	// 	return nil, err
	// }

	content, err := service.expermentgeminiHandler.GetExperimentContent(experiment.ExperimentName, promptMessage)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (service *ExperimentService) GenerateExperimentsPerChapter(generateDTO *dtos.GenerateExperimentDTO) ([]domain.Experiment, error) {
	// Get the experiments from Gemini
	experiments, err := service.expermentgeminiHandler.GetExperimentsForChapter(generateDTO, generateDTO.ChapterName)
	if err != nil {
		return nil, err
	}

	// Save the experiments to the database and return them
	domainExperiments := []domain.Experiment{}
	for _, exp := range experiments {
		exp, err := service.experimentRepository.Create(&exp)
		if err != nil {
			return nil, err
		}
		domainExperiments = append(domainExperiments, *exp)
	}

	return domainExperiments, nil
}

func (service *ExperimentService) CreateExperiment(experimentDTO *dtos.ExperimentDTO) (*domain.Experiment, error) {
	experiment := &domain.Experiment{
		ExperimentName: experimentDTO.ExperimentName,
		SubjectID:      experimentDTO.SubjectID,
		ChapterID:      experimentDTO.ChapterID,
	}
	exp, err := service.experimentRepository.Create(experiment)
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func (service *ExperimentService) UpdateExperiment(experimentID uint, experimentDTO *dtos.ExperimentDTO) (*domain.Experiment, error) {
	experiment, err := service.experimentRepository.GetById(experimentID)

	if err != nil {
		return nil, err
	}

	experiment.ExperimentName = experimentDTO.ExperimentName
	exp, err := service.experimentRepository.Update(experiment)
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func (service *ExperimentService) DeleteExperiment(experimentID uint) error {
	experiment, err := service.experimentRepository.GetById(experimentID)
	if err != nil {
		return err
	}

	if _, err := service.experimentRepository.Delete(experiment); err != nil {
		return err
	}

	return nil
}
