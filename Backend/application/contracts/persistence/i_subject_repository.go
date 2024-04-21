package contracts

import (
	"Brilliant/domain"
)

type ISubjectRepository interface {
	IGenericRepository[domain.Subject]
	CreateSubject(SubjectName string, UserID uint) (*domain.Subject, error)
	SearchSubjectsByName(query string) ([]*domain.Subject, error)
	GenerateRoadMap(subject *domain.Subject) ([]*domain.Chapter, []*domain.Lesson, error)
}
