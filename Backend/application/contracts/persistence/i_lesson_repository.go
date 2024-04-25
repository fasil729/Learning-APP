package contracts

import "Brilliant/domain"

type ILessonRepository interface {
	IGenericRepository[domain.Lesson]
	CreateLesson(chapterID uint, lessonName string) (*domain.Lesson, error)
	GetLessonsByChapterId(chapterId uint) ([]*domain.Lesson, error)
}
