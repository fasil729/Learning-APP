package dtos

// AddNoteDTO represents the data required to add a note to a chapter
type AddNoteDTO struct {
	Text      string // Text content of the note
	ImageData []byte // Image data to be added
}
