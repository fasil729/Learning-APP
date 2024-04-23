package contracts

import "Brilliant/domain"

type ILessonRepository interface {
	IGenericRepository[domain.Lesson]
	CreateLesson(lessonName string) (*domain.Lesson, error)
	GetLessonsByChapterId(chapterId uint) ([]*domain.Lesson, error)
}
