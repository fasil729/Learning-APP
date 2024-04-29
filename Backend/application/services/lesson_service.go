package services

import (
	contracts "Brilliant/application/contracts/persistence"
	contract "Brilliant/application/contracts/gemini"
	dtos "Brilliant/application/dtos/lesson"
	"Brilliant/domain"
	"fmt"
	// "os"
)

type LessonService struct {
	LessonRepository contracts.ILessonRepository
	subjectGeminiHandler contract.IGeminiSubjectHandler
}

func NewLessonService(lessonRepository contracts.ILessonRepository, subjectGeminiHandler contract.IGeminiSubjectHandler) *LessonService {
	return &LessonService{
		LessonRepository: lessonRepository,
		subjectGeminiHandler: subjectGeminiHandler,
		
	}
}

func (service *LessonService) CreateLesson(lessonDTO *dtos.LessonDTO) (*domain.Lesson, error) {
	lesson := &domain.Lesson{
		Name:      lessonDTO.LessonName,
		ChapterID: lessonDTO.ChapterID,
	}

	lesson, err := service.LessonRepository.Create(lesson)
	if err != nil {
		return nil, err
	}
	return lesson, nil

}

func (service *LessonService) UpdateLesson(lessonId uint, lessonDTO *dtos.LessonDTO) (*domain.Lesson, error) {
	lesson, err := service.LessonRepository.GetById(lessonId)
	if err != nil {
		return nil, err
	}

	lesson.Name = lessonDTO.LessonName

	lesson, err = service.LessonRepository.Update(lesson)
	if err != nil {
		return nil, err
	}
	return lesson, nil
}

func (service *LessonService) DeleteLesson(lessonId uint) error {
	lesson, err := service.LessonRepository.GetById(lessonId)
	if err != nil {
		return err
	}

	_, err = service.LessonRepository.Delete(lesson)
	if err != nil {
		return err
	}
	return nil
}

// GetLessonContent retrieves content from a file, simulating an API call in the future.
func (service *LessonService) GetLessonContent(LessonID uint) ([]byte, error) {
	// Fetch the lesson from the repository
	lesson, err := service.LessonRepository.GetById(LessonID)
	if err != nil {
		return nil, fmt.Errorf("failed to get lesson: %w", err)
	}

	// // If there's no content link, set a dummy path
	// if lesson.ContentLink == "" {
	// 	lesson.ContentLink = "data/lesson/lessonContent.md" // Default content path
	// 	_, err := service.LessonRepository.Update(lesson)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to update lesson with content link: %w", err)
	// 	}
	// }

	// // Verify if the file exists
	// if _, err := os.Stat(lesson.ContentLink); os.IsNotExist(err) {
	// 	return nil, fmt.Errorf("content not found at path: %s", lesson.ContentLink)
	// }

	// // Read the content from the file
	// content, err := os.ReadFile(lesson.ContentLink)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read content file: %w", err)
	// }
    content, err := service.subjectGeminiHandler.GenerateLessonDetailContent(lesson.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to generate lesson content: %w", err)
	}
	return content, nil
}

func (service *LessonService) GenerateLessons(chapterID uint) ([]*domain.Lesson, error) {
	lessonNames := []string{
		"Introduction to the Nervous System",
		"Structure of the Nervous System",
		"Function of the Nervous System",
		"Visualization of the Nervous System",
		"Introduction",
		"Learning Objectives",
		"Conclusion",
	}

	var lessons []*domain.Lesson
	for _, lessonName := range lessonNames {
		lesson := domain.Lesson{
			Name:      lessonName,
			ChapterID: chapterID,
		}

		createdLesson, err := service.LessonRepository.Create(&lesson)
		if err != nil {
			return nil, err
		}

		lessons = append(lessons, createdLesson)
	}

	return lessons, nil
}

// 	// Generate content for the lesson

// 	// Create a new Markdown file and write the content to it
// 	contentFile, err := os.Create(filepath.Join("lessons", strconv.FormatUint(uint64(lesson.ID), 10)+".md"))
// 	if err != nil {
// 		return nil, err

// 	}
// 	defer contentFile.Close()

// 	_, err = contentFile.WriteString(content)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Store the link to the Markdown file in the database
// 	lesson.ContentLink = contentFile.Name()

// 	response, err := repository.Create(&lesson)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return response, nil
