package services

import (
	"Brilliant/application/contracts/persistence"
	"Brilliant/application/dtos/experiment"
	"Brilliant/domain"
	"io/ioutil"
)

type ExperimentService struct {
	experimentRepository contracts.IExperimentRepository
}

func NewExperimentService(experimentRepository contracts.IExperimentRepository) *ExperimentService {
	return &ExperimentService{
		experimentRepository: experimentRepository,
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

	if experiment.ContentLink == "" {
		// Gemini API call will be done here

		// For now, return a dummy experiment
		experiment.ContentLink = "data/experiment/generated_experiment.json"

		// Update the experiment with the content link
		if _, err := service.experimentRepository.Update(experiment); err != nil {
			return nil, err
		}
	}

	// Read the content of the JSON file
	content, err := ioutil.ReadFile(experiment.ContentLink)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (service *ExperimentService) GenerateExperimentsPerChapter(generateDTO *dtos.GenerateExperimentDTO) ([]*domain.Experiment, error) {
	// For now, return three dummy experiments
	experiments := []*domain.Experiment{
		{ExperimentName: "Experiment 1", SubjectID: generateDTO.SubjectID, ChapterID: generateDTO.ChapterID},
		{ExperimentName: "Experiment 2", SubjectID: generateDTO.SubjectID, ChapterID: generateDTO.ChapterID},
		{ExperimentName: "Experiment 3", SubjectID: generateDTO.SubjectID, ChapterID: generateDTO.ChapterID},
	}

	for _, exp := range experiments {
		service.experimentRepository.Create(exp)
	}

	return experiments, nil
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



