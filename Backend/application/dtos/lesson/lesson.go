package dtos

type LessonDTO struct {
	ChapterID   uint   `json:"id"`
	LessonName  string `json:"chapterName"`
	ContentLink string `json:"contentLink"`
}
