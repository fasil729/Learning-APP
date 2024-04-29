package contracts

import "Brilliant/domain"

type INoteRepository interface {
	IGenericRepository[domain.Note]
	CreateNote(chapterID uint, ContentLink string) (*domain.Note, error)
	GetNotesByChapterID(chapterID uint) (*domain.Note, error)
}
