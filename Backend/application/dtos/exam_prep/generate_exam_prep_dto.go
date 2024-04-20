package dtos

import (
	"time"
)

type GenerateExamPrepDTO struct {
	Topics   []string
	Prompt   string
	ReadTime time.Duration
}
