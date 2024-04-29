package contracts


// Define the ISubjectHandler interface in the contracts package
type IGeminiSubjectHandler interface {
	GenerateTableOfContent(subjectName string) ([]Chapter, error)
	GenerateLessonDetailContent(lessonName string) ([]byte, error)
}

type Chapter struct {
	Name    string   `json:"name"`
	Lessons []string `json:"lessons"`
}
