package repositories

import (
	contracts "Brilliant/application/contracts/persistence"
	"Brilliant/domain"

	"gorm.io/gorm"
)

type ChapterRepository struct {
	*GenericRepository[domain.Chapter]
	db *gorm.DB
}

func NewChapterRepository(db func() *gorm.DB) contracts.IChapterRepository {
	return &ChapterRepository{
		db: db(),
	}
}

func (repository *ChapterRepository) CreateChapter(chapter *domain.Chapter) (*domain.Chapter, error) {
	createdChapter := &domain.Chapter{
		ChapterName: chapter.ChapterName,
	}
	err := repository.db.Create(createdChapter).Error
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
	err := repository.db.Delete(chapter).Error
	if err != nil {
		return err
	}

	return nil
}
