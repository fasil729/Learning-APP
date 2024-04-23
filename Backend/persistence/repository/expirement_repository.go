package repositories

import (
	contracts "Brilliant/application/contracts/persistence"
	"Brilliant/domain"

	"gorm.io/gorm"
)

type ExperimentRepository struct {
	GenericRepository[domain.Experiment]
	database *gorm.DB
	
}

func NewExperimentRepository(getDb func() *gorm.DB) contracts.IExperimentRepository  {
	return &ExperimentRepository{
		database: getDb(),
	}
}
func (r *ExperimentRepository) GetByChapterID(chapterID uint) ([]*domain.Experiment, error) {
	var experiments []*domain.Experiment
	if err := r.database.Where("chapter_id = ?", chapterID).Find(&experiments).Error; err != nil {
		return nil, err
	}
	return experiments, nil
}

func (r *ExperimentRepository) GetBySubjectID(subjectID uint) ([]*domain.Experiment, error) {
	var experiments []*domain.Experiment
	if err := r.database.Where("subject_id = ?", subjectID).Find(&experiments).Error; err != nil {
		return nil, err
	}
	return experiments, nil
}