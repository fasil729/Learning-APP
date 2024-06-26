package contracts

import (
	"Brilliant/domain"
)

type ISubjectRepository interface {
	IGenericRepository[domain.Subject]
	CreateSubject(SubjectName string, UserID uint) (*domain.Subject, error)
	SearchSubjectsByName(userId uint, query string) ([]*domain.Subject, error)
}
