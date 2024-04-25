package contracts

import "github.com/google/generative-ai-go/genai"

// IExperimentHandler defines the contract for handling experiments in Gemini.
type IExperimentHandler interface {
	GetExperimentsForChapter(chapterName string) ([]*genai.Candidate, error)
	GetExperimentContent(experimentName string) ([]*genai.Candidate, error)
}