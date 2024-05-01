package services

import (
	contract "Brilliant/application/contracts/gemini"
	contracts "Brilliant/application/contracts/persistence"
	dtos "Brilliant/application/dtos/note"

	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type NoteService struct {
	NoteRepository    contracts.INoteRepository
	ChapterRepository contracts.IChapterRepository
	NotegeminiHandler contract.INoteHandler
}

func NewNoteService(noteRepository contracts.INoteRepository, chapterRepository contracts.IChapterRepository, notegeminiHandler contract.INoteHandler) *NoteService {
	return &NoteService{
		NoteRepository:    noteRepository,
		ChapterRepository: chapterRepository,
		NotegeminiHandler: notegeminiHandler,
	}
}

func (service *NoteService) AddNote(chapterID uint, request *dtos.AddNoteDTO) ([]byte, error) {
	// Get the chapter by ID
	chapter, err := service.ChapterRepository.GetChapterByID(chapterID)
	fmt.Printf("chapter: %v", chapter)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chapter by ID")
	}

	// Check if the content link exists
	contentLink := chapter.NoteLink
	contentFolder := fmt.Sprintf("data/chapter/%d/content", chapterID)

	if contentLink == "" {
		// Create a new Markdown file for the chapter
		if err := os.MkdirAll(contentFolder, 0755); err != nil {
			return nil, errors.Wrap(err, "failed to create content folder")
		}
		contentLink = filepath.Join(contentFolder, fmt.Sprintf("chapter_%d.md", chapterID))
		if _, err := os.Create(contentLink); err != nil {
			return nil, errors.Wrap(err, "failed to create content file")
		}

		chapter.NoteLink = contentLink
		service.NoteRepository.CreateNote(chapterID, contentLink)

		if request.ImageData != nil {
			newContent, err := service.NotegeminiHandler.AddNoteForChapterFromImage(request.Text, request.ImageData, chapter.Name)
			if err != nil {
				return nil, errors.Wrap(err, "failed to generate note from image")
			}

			err = os.WriteFile(contentLink, newContent, 0644)
			if err != nil {
				return nil, errors.Wrap(err, "failed to write text to file")
			}
		} else {
			err = os.WriteFile(contentLink, []byte(request.Text), 0644)
			if err != nil {
				return nil, errors.Wrap(err, "failed to write text to file")
			}
		}

	} else {
		// Append the note to the chapter content
		Previouscontent, err := os.ReadFile(contentLink)

		if err != nil {
			return nil, errors.Wrap(err, "failed to read content file")
		}

		file, err := os.OpenFile(contentLink, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, errors.Wrap(err, "failed to open content file")
		}
		defer file.Close()

		newContent, err := service.NotegeminiHandler.AddNoteForChapter(string(Previouscontent), request.Text, chapter.Name)

		if err != nil {
			return nil, errors.Wrap(err, "failed to generate note")
		}

		if request.ImageData != nil {
			newContentImage, err := service.NotegeminiHandler.AddNoteForChapterFromImage(string(newContent), request.ImageData, chapter.Name)
			if err != nil {
				return nil, errors.Wrap(err, "failed to generate note from image")
			}

			err = os.WriteFile(contentLink, newContentImage, 0644)
			if err != nil {
				return nil, errors.Wrap(err, "failed to write text to file")
			}

		} else {
			_, err = file.WriteString(string(newContent) + "\n")
			if err != nil {
				return nil, errors.Wrap(err, "failed to write text to file")
			}

		}

	}
	content, err := os.ReadFile(contentLink)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read content file")
	}

	return content, nil
}

func (service *NoteService) GetNoteByChapterID(chapterID uint) ([]byte, error) {
	// Get the chapter by ID
	chapter, err := service.ChapterRepository.GetChapterByID(chapterID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chapter by ID")
	}

	// Check if the content link exists
	contentLink := chapter.NoteLink
	if contentLink == "" {
		return nil, errors.New("no notes found for the chapter")
	}

	// Read the file and return its content
	content, err := os.ReadFile(contentLink)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read content file")
	}

	return content, nil
}
