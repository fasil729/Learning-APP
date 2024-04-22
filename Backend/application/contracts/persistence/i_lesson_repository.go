package contracts

import "Brilliant/domain"

type ILessonRepository interface {
	IGenericRepository[domain.Lesson]
	CreateLesson(name string, chapterId uint) (*domain.Lesson, error)
}
