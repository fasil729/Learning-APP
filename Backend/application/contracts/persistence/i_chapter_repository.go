package contracts

import (
	"Brilliant/domain"
)

type IChapterRepository interface {
	IGenericRepository[domain.Chapter]
	CreateChapter(chapter *domain.Chapter) (*domain.Chapter, error)
}
