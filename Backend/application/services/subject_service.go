package services

import (
	"hackathon.com/leariningApp/application/contracts"
	"hackathon.com/leariningApp/domain"
)

type SubjectService struct {
	subjectRepository contracts.ISubjectRepository
}

func NewSubjectService(subjectRepository contracts.ISubjectRepository) *SubjectService {
	return &SubjectService{subjectRepository: subjectRepository}
}

func (service *SubjectService) CreateSubject(SubjectName string, UserID uint) (*domain.Subject, error) {
	subject, err := service.subjectRepository.CreateSubject(SubjectName, UserID)
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

func (service *SubjectService) GenerateRoadMap(subject *domain.Subject) ([]*domain.Chapter, []*domain.Lesson, error) {
	chapters, lessons, err := service.subjectRepository.GenerateRoadMap(subject)
	if err != nil {
		return nil, nil, err
	}
	return chapters, lessons, nil
}
