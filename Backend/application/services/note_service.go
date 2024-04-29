package services

import (
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
}

func NewNoteService(noteRepository contracts.INoteRepository, chapterRepository contracts.IChapterRepository) *NoteService {
	return &NoteService{
		NoteRepository:    noteRepository,
		ChapterRepository: chapterRepository,
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
		service.ChapterRepository.Update(chapter)
	}

	// Write the text to the file
	err = os.WriteFile(contentLink, []byte(request.Text+"\n"), 0644)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write text to file")
	}

	// If there's an image, save the image data to a file and reference it in the Markdown file
	if len(request.ImageData) > 0 {
		imagePath := filepath.Join(contentFolder, "image.jpg")
		err = os.WriteFile(imagePath, request.ImageData, 0644)
		if err != nil {
			return nil, errors.Wrap(err, "failed to write image to file")
		}

		file, err := os.OpenFile(contentLink, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return nil, errors.Wrap(err, "failed to open file")
		}
		defer file.Close()

		_, err = file.WriteString("\n\n![Image](" + imagePath + ")\n")
		if err != nil {
			return nil, errors.Wrap(err, "failed to write image reference to file")
		}
	}

	// Read the file back and return its content
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
