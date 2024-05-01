package contracts

// INoteHandler defines the contract for handling notes in Gemini.
type INoteHandler interface {
	AddNoteForChapterFromImage(previousContent string, imageData []byte, chapterName string) ([]byte, error)
	AddNoteForChapter(previousContent, noteContent, chapterName string) ([]byte, error)
}
