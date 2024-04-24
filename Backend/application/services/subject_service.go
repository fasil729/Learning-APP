package services

import (
	contracts "Brilliant/application/contracts/persistence"
	Chapterdtos "Brilliant/application/dtos/chapter"
	Lessondtos "Brilliant/application/dtos/lesson"
	Subjctdtos "Brilliant/application/dtos/subject"

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

func (service *SubjectService) CreateSubject(UserId uint, subjectDTO *Subjctdtos.CreateSubjectDTO) (*Subjctdtos.SubjectDTO, error) {
	subject, err := service.subjectRepository.CreateSubject(subjectDTO.SubjectName, UserId)
	if err != nil {
		return nil, err
	}

	// Dummy data for chapters and lessons
	chapters := []domain.Chapter{
		{
			Name: "Chapter 1",
			Lessons: []domain.Lesson{
				{Name: "Lesson 1.1"},
				{Name: "Lesson 1.2"},
			},
		},
		{
			Name: "Chapter 2",
			Lessons: []domain.Lesson{
				{Name: "Lesson 2.1"},
				{Name: "Lesson 2.2"},
			},
		},
	}

	chaptersDTO := make([]Chapterdtos.ChapterDTO, len(chapters))

	for i, chapter := range chapters {
		lessonDTO := make([]Lessondtos.LessonDTO, len(chapter.Lessons))

		for j, lesson := range chapter.Lessons {
			lessonDTO[j] = Lessondtos.LessonDTO{
				LessonName: lesson.Name,
			}
		}
		chaptersDTO[i] = Chapterdtos.ChapterDTO{
			ChapterName: chapter.Name,
			Lessons:     lessonDTO,
		}
	}

	for i, chapterDTO := range chapters {
		chapter, err := service.chapterRepository.CreateChapter(chapterDTO.Name, subject.ID)
		if err != nil {
			return nil, err
		}

		chapterDTO.ID = chapter.ID

		for j, lessonDTO := range chapterDTO.Lessons {
			lesson, err := service.lessonRepository.CreateLesson(chapter.ID, lessonDTO.Name)
			if err != nil {
				return nil, err
			}

			chapterDTO.Lessons[j].ID = lesson.ID
		}

		chapters[i] = chapterDTO
	}

	subjectDto := &Subjctdtos.SubjectDTO{
		ID:                 subject.ID,
		UserID:             subject.UserID,
		SubjectName:        subject.Name,
		Chapters:           chaptersDTO,
		TextBookLink:       subject.TextBookLink,
		ReferenceBooksLink: "",
	}

	return subjectDto, nil
}

func (service *SubjectService) SearchSubjectsByName(query string) ([]*domain.Subject, error) {
	subjects, err := service.subjectRepository.SearchSubjectsByName(query)
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
