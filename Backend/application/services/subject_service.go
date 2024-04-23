package services

import (
	"Brilliant/application/contracts/persistence"
	dtos "Brilliant/application/dtos/subject"
	// dtos "Brilliant/application/dtos/subject"
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

func (service *SubjectService) CreateSubject(subjectDTO *dtos.SubjectDTO) (*domain.Subject, error) {
	subject := &domain.Subject{
		SubjectName: subjectDTO.SubjectName,
		UserID:      subjectDTO.UserID,
	}

	subject, err := service.subjectRepository.CreateSubject(subject.SubjectName, subject.UserID)
	if err != nil {
		return nil, err
	}
	return subject, nil
}

func (service *SubjectService) SearchSubjectsByName(query string) ([]*domain.Subject, error) {
	subjects, err := service.subjectRepository.SearchSubjectsByName(query)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}

func (service *SubjectService) GenerateRoadMap(GenerateSubjectDTO *dtos.GenerateSubjectDTO) ([]*domain.Chapter, []*domain.Lesson, error) {
	// Create dummy chapters
	chapters := []*domain.Chapter{
		{ChapterName: "Chapter 1", SubjectID: GenerateSubjectDTO.SubjectID},
		{ChapterName: "Chapter 2", SubjectID: GenerateSubjectDTO.SubjectID},
		// Add more chapters as needed
	}
	lessons := []*domain.Lesson{
		{LessonName: "Lesson 1", ChapterID: GenerateSubjectDTO.ChapterID},
		{LessonName: "Lesson 2", ChapterID: GenerateSubjectDTO.ChapterID},
		// Add more lessons as needed
	}

	for _, chapter := range chapters {
		// Save the chapter to the database
		var err error
		_, err = service.chapterRepository.CreateChapter(chapter)
		if err != nil {
			return nil, nil, err
		}

		// Create dummy lessons for the chapter

		for _, lesson := range lessons {
			// Save the lesson to the database
			_, err = service.lessonRepository.CreateLesson(lesson.LessonName)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return chapters, lessons, nil
}
