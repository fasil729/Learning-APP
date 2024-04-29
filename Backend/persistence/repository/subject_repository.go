package repositories

import (
	contracts "Brilliant/application/contracts/persistence"
	"Brilliant/domain"
	"Brilliant/persistence/migrations"

	"gorm.io/gorm"
	// "log"
)

type SubjectRepository struct {
	GenericRepository[domain.Subject]
	database *gorm.DB
}

func NewSubjectRepository(getDb func() *gorm.DB) contracts.ISubjectRepository {
	db := getDb()
	return &SubjectRepository{
		GenericRepository: GenericRepository[domain.Subject]{database: db},
		database:          db,
	}
}

func (repository *SubjectRepository) CreateSubject(SubjectName string, UserID uint) (*domain.Subject, error) {
	subject := domain.Subject{
		Name:   SubjectName,
		UserID: UserID,
	}

	createdSubject, err := repository.Create(&subject)

	if err != nil {
		return nil, err
	}

	return createdSubject, nil
}

func (repository *SubjectRepository) UpdateSubject(subject *domain.Subject) (*domain.Subject, error) {
	updatedSubject, err := repository.Update(subject)

	if err != nil {
		return nil, err
	}

	return updatedSubject, nil
}

func (repository *SubjectRepository) DeleteSubject(subject *domain.Subject) error {
	_, err := repository.Delete(subject)

	if err != nil {
		return err
	}

	return nil
}

func (repository *SubjectRepository) SearchSubjectsByName(UserID uint, query string) ([]*domain.Subject, error) {
	var searchedSubjects []*domain.Subject

	response := migrations.GetDb().Where("user_id = ? AND Name LIKE ?", UserID, "%"+query+"%").Find(&searchedSubjects)
	error := response.Error

	if error != nil {
		return nil, error
	}

	return searchedSubjects, nil

}

// func (repository *SubjectRepository) CreateChapterAndLesson(subject *domain.Subject, chapterName string, lessonName string, contentLink string) (*domain.Chapter, *domain.Lesson, error) {
// 	// Create a new chapter
// 	chapter := &domain.Chapter{
// 		Name:      chapterName,
// 		SubjectID: subject.ID,
// 	}

// 	// Create a new lesson
// 	lesson := &domain.Lesson{
// 		LessonName:  lessonName,
// 		ChapterID:   chapter.ID,
// 		ContentLink: contentLink,
// 	}

// 	// Save the chapter and lesson to the database
// 	if err := repository.database.Create(chapter).Error; err != nil {
// 		return nil, nil, err
// 	}
// 	if err := repository.database.Create(lesson).Error; err != nil {
// 		return nil, nil, err
// 	}

// 	return chapter, lesson, nil
// }

// func (repository *SubjectRepository) GenerateRoadMap(subject *domain.Subject) ([]*domain.Chapter, []*domain.Lesson, error) {
// 	// Generate dummy chapters and lessons for the subject
// 	chapters := []*domain.Chapter{
// 		{Name: "Chapter 1", SubjectID: subject.ID},
// 		{Name: "Chapter 2", SubjectID: subject.ID},
// 		// Add more chapters as needed
// 	}

// 	lessons := []*domain.Lesson{
// 		{Name: "Lesson 1", ChapterID: chapters[0].ID, ContentLink: "https://example.com/lesson1"},
// 		{Name: "Lesson 2", ChapterID: chapters[0].ID, ContentLink: "https://example.com/lesson2"},
// 		// Add more lessons as needed
// 	}

// 	// Save the generated chapters and lessons to the database
// 	if err := repository.database.Create(&chapters).Error; err != nil {
// 		return nil, nil, err
// 	}
// 	if err := repository.database.Create(&lessons).Error; err != nil {
// 		return nil, nil, err
// 	}

// 	return chapters, lessons, nil
// }
