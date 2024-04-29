package contracts


// INoteHandler defines the contract for handling notes in Gemini.
type INoteHandler interface {
	AddNoteForChapterFromImage(imageData []byte, chapterName string) error
	AddNoteForChapter(noteContent, chapterName string) error
}