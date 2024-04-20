package contracts

import "Brilliant/domain"

type IExperimentRepository interface {
	IGenericRepository[domain.Experiment]
	GetByChapterID(chapterID uint) ([]*domain.Experiment, error)
	GetBySubjectID(topicID uint) ([]*domain.Experiment, error)
}