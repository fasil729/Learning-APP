package dtos

// ExperimentDTO represents the Experiment Data Transfer Object
type ExperimentDTO struct {
	ExperimentName string `json:"experiment_name"`
	SubjectID        uint   `json:"topic_id"`
	ChapterID      uint   `json:"chapter_id"`
	ContentLink  string `json:"content_link"`
}

