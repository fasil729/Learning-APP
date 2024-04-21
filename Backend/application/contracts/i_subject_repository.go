package contracts

import (
	"hackathon.com/leariningApp/domain"
)

type ISubjectRepository interface {
	IGenericRepository[domain.Subject]
	CreateSubject(SubjectName string, UserID uint) (*domain.Subject, error)
	SearchSubjectsByName(query string) ([]*domain.Subject, error)
	GenerateRoadMap(subject *domain.Subject) ([]*domain.Chapter, []*domain.Lesson, error)
}
