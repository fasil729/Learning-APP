package services

import (
	contracts "Brilliant/application/contracts/persistence"
	dtos "Brilliant/application/dtos/subject"
	// Chapterdtos "Brilliant/application/dtos/chapter"
	// Lessondtos "Brilliant/application/dtos/lesson"
	// Subjctdtos "Brilliant/application/dtos/subject"

	"Brilliant/domain"
)

type SubjectService struct {
	subjectRepository contracts.ISubjectRepository
	chapterRepository contracts.IChapterRepository
	lessonRepository  contracts.ILessonRepository
}

func NewSubjectService(subjectRepository contracts.ISubjectRepository, chapterRepository contracts.IChapterRepository, lessonRepository contracts.ILessonRepository) *SubjectService {
	return &SubjectService{
		subjectRepository: subjectRepository,
		chapterRepository: chapterRepository,
		lessonRepository:  lessonRepository}
}

// CreateSubject creates a new subject along with its chapters and lessons.
func (service *SubjectService) CreateSubject(userID uint, SubjectDTO *dtos.CreateSubjectDTO) (*domain.Subject, error) {
	// Create the subject in the repository
	createdSubject, err := service.subjectRepository.CreateSubject(SubjectDTO.SubjectName, userID)
	if err != nil {
		return nil, err
	}

	// Dummy data for chapters and lessons
	chapters := []struct {
		Name    string
		Lessons []string
	}{
		{
			Name: "Chapter 1",
			Lessons: []string{
				"Lesson 1.1",
				"Lesson 1.2",
			},
		},
		{
			Name: "Chapter 2",
			Lessons: []string{
				"Lesson 2.1",
				"Lesson 2.2",
			},
		},
	}

	var domainChapters []domain.Chapter

	// Create chapters and their lessons in the database
	for _, chapterData := range chapters {
		createdChapter, err := service.chapterRepository.CreateChapter(chapterData.Name, createdSubject.ID)
		if err != nil {
			return nil, err
		}

		var domainLessons []domain.Lesson
		for _, lessonName := range chapterData.Lessons {
			createdLesson, err := service.lessonRepository.CreateLesson(createdChapter.ID, lessonName)
			if err != nil {
				return nil, err
			}

			domainLessons = append(domainLessons, *createdLesson) // Adding created lessons with IDs
		}

		createdChapter.Lessons = domainLessons                   // Assigning lessons with correct IDs
		domainChapters = append(domainChapters, *createdChapter) // Adding created chapters with IDs
	}

	createdSubject.Chapters = domainChapters // Assigning chapters with correct IDs

	return createdSubject, nil
}

func (service *SubjectService) SearchSubjectsByName(UserID uint, query string) ([]*domain.Subject, error) {
	subjects, err := service.subjectRepository.SearchSubjectsByName(UserID, query)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}

// func (service *SubjectService) GenerateRoadMap(GenerateSubjectDTO *Subjctdtos.GenerateSubjectDTO) ([]*domain.Chapter, []*domain.Lesson, error) {
// 	// Create dummy chapters
// 	chapters := []*domain.Chapter{
// 		{ChapterName: "Chapter 1", SubjectID: GenerateSubjectDTO.SubjectID},
// 		{ChapterName: "Chapter 2", SubjectID: GenerateSubjectDTO.SubjectID},
// 		// Add more chapters as needed
// 	}
// 	lessons := []*domain.Lesson{
// 		{Name: "Lesson 1", ChapterID: GenerateSubjectDTO.ChapterID},
// 		{LessonName: "Lesson 2", ChapterID: GenerateSubjectDTO.ChapterID},
// 		// Add more lessons as needed
// 	}

// 	for _, chapter := range chapters {
// 		// Save the chapter to the database
// 		var err error
// 		_, err = service.chapterRepository.CreateChapter(chapter)
// 		if err != nil {
// 			return nil, nil, err
// 		}

// 		// Create dummy lessons for the chapter

// 		for _, lesson := range lessons {
// 			// Save the lesson to the database
// 			_, err = service.lessonRepository.CreateLesson(lesson.LessonName)
// 			if err != nil {
// 				return nil, nil, err
// 			}
// 		}
// 	}

// 	return chapters, lessons, nil
// }
