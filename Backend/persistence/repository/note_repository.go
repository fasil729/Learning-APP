package repositories

import (
	contracts "Brilliant/application/contracts/persistence"
	"Brilliant/domain"
	"fmt"
	"gorm.io/gorm"
)

type NoteRepository struct {
	GenericRepository[domain.Note]
	database *gorm.DB
}

func NewNoteRepository(getDb func() *gorm.DB) contracts.INoteRepository {
	db := getDb()
	return &NoteRepository{
		GenericRepository: GenericRepository[domain.Note]{database: db},
		database:          db,
	}
}

func (repository *NoteRepository) CreateNote(chapterID uint, noteLink string) (*domain.Note, error) {
	note := &domain.Note{
		ChapterID:   chapterID,
		ContentLink: noteLink,
	}
	createdNoteLink, err := repository.Create(note)
	if err != nil {
		return nil, err
	}
	return createdNoteLink, nil

}

func (repository *NoteRepository) GetNotesByChapterID(chapterID uint) (*domain.Note, error) {
	fmt.Printf("chapterID: %v", chapterID)
	var note *domain.Note
	if err := repository.database.Where("chapter_id = ?", chapterID).Find(&note).Error; err != nil {
		return nil, err
	}
	return note, nil
}
