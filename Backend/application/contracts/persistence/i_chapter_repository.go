package contracts

import (
	"Brilliant/domain"
)

type IChapterRepository interface {
	IGenericRepository[domain.Chapter]
	CreateChapter(chapterName string, subjectID uint) (*domain.Chapter, error)
}
