package repositories

import (
	contracts "Brilliant/application/contracts/persistence"
	"Brilliant/domain"

	"gorm.io/gorm"
)

type ChapterRepository struct {
	GenericRepository[domain.Chapter]
	database *gorm.DB
}

func NewChapterRepository(getDb func() *gorm.DB) contracts.IChapterRepository {
	db := getDb()
	return &ChapterRepository{
		GenericRepository: GenericRepository[domain.Chapter]{database: db},
		database:          db,
	}
}

func (repository *ChapterRepository) CreateChapter(chapterName string, subjectID uint) (*domain.Chapter, error) {
	chapter := &domain.Chapter{
		Name:      chapterName,
		SubjectID: subjectID,
	}
	createdChapter, err := repository.Create(chapter)
	if err != nil {
		return nil, err
	}

	return createdChapter, nil
}

func (repository *ChapterRepository) UpdateChapter(chapter *domain.Chapter) (*domain.Chapter, error) {
	updated, err := repository.Update(chapter)

	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (repository *ChapterRepository) DeleteChapter(chapter *domain.Chapter) error {
	_, err := repository.Delete(chapter)
	if err != nil {
		return err
	}

	return nil
}
