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

// type UserRepository struct {
// 	*GenericRepository[domain.User]
// }

func NewSubjectRepository(getDb func() *gorm.DB) contracts.ISubjectRepository {
	db := getDb()
	return &SubjectRepository{
		GenericRepository: GenericRepository[domain.Subject]{database: db},
		database:          db,
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

func (repository *SubjectRepository) SearchSubjectsByName(query string) ([]*domain.Subject, error) {
	var searchedTopics []*domain.Subject

	response := migrations.GetDb().Where("TopicName LIKE ?", "%"+query+"%").Find(&searchedTopics)
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
