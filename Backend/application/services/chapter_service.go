package services

import (
	contracts "Brilliant/application/contracts/persistence"
	dtos "Brilliant/application/dtos/chapter"
	"Brilliant/domain"
)

type ChapterService struct {
	ChapterRepository contracts.IChapterRepository
}

func NewChapterService(chapterRepository contracts.IChapterRepository) *ChapterService {
	return &ChapterService{
		ChapterRepository: chapterRepository,
	}
}

func (service *ChapterService) CreateChapter(chapterDTO *dtos.ChapterDTO) (*domain.Chapter, error) {
	chapter := &domain.Chapter{
		ChapterName: chapterDTO.ChapterName,
		SubjectID:   chapterDTO.SubjectID,
	}

	chapter, err := service.ChapterRepository.Create(chapter)
	if err != nil {
		return nil, err
	}
	return chapter, nil
}

func (service *ChapterService) UpdateChapter(chapterId uint, chapterDTO *dtos.ChapterDTO) (*domain.Chapter, error) {
	chapter, err := service.ChapterRepository.GetById(chapterId)
	if err != nil {
		return nil, err
	}

	chapter.ChapterName = chapterDTO.ChapterName

	chapter, err = service.ChapterRepository.Update(chapter)
	if err != nil {
		return nil, err
	}
	return chapter, nil
}

func (service *ChapterService) DeleteChapter(chapterId uint) error {
	chapter, err := service.ChapterRepository.GetById(chapterId)
	if err != nil {
		return err
	}

	_, err = service.ChapterRepository.Delete(chapter)
	if err != nil {
		return err
	}
	return nil
}
