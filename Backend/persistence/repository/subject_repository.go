package repositories

import (
	"gorm.io/gorm"
	"Brilliant/domain"
	"Brilliant/persistence/migrations"
	// "log"
)

type SubjectRepository struct {
	*GenericRepository[domain.Subject]
}

type UserRepository struct {
	*GenericRepository[domain.User]
}

func NewSubjectRepository(GetDb func() *gorm.DB) *SubjectRepository {
	return &SubjectRepository{
		GenericRepository: NewGenericRepository[domain.Subject](GetDb),
	}
}

func (repository *SubjectRepository) CreateSubject(SubjectName string, UserID uint) (*domain.Subject, error) {
	subject := domain.Subject{
		SubjectName: SubjectName,
		UserID:      UserID,
	}

	createdSubject, err := repository.Create(&subject)

	if err != nil {
		return nil, err
	}

	return createdSubject, nil
}

func searchSubjectsByName(query string) ([]*domain.Subject, error) {
	var searchedTopics []*domain.Subject

	response := migrations.GetDB().Where("TopicName LIKE ?", "%"+query+"%").Find(&searchedTopics)
	error := response.Error

	if error != nil {
		return nil, error
	}

	return searchedTopics, nil

}

func (repository *SubjectRepository) GenerateRoadMap(subject *domain.Subject) ([]*domain.Chapter, []*domain.Lesson, error) {
	// Generate dummy chapters and lessons for the subject
	chapters := []*domain.Chapter{
		{Name: "Chapter 1", SubjectID: subject.ID},
		{Name: "Chapter 2", SubjectID: subject.ID},
		// Add more chapters as needed
	}

	lessons := []*domain.Lesson{
		{Name: "Lesson 1", ChapterID: chapters[0].ID, ContentLink: "https://example.com/lesson1"},
		{Name: "Lesson 2", ChapterID: chapters[0].ID, ContentLink: "https://example.com/lesson2"},
		// Add more lessons as needed
	}

	// Save the generated chapters and lessons to the database
	if err := repository.database.Create(&chapters).Error; err != nil {
		return nil, nil, err
	}
	if err := repository.database.Create(&lessons).Error; err != nil {
		return nil, nil, err
	}

	return chapters, lessons, nil
}
