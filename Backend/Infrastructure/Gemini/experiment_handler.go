package gemini

import (
	"context"
	"encoding/json"
	"log"

	"Brilliant/application/contracts/gemini"
	dtos "Brilliant/application/dtos/experiment"
	"Brilliant/domain"

	"github.com/google/generative-ai-go/genai"
)

// GeminiExperimentHandler implements IExperimentHandler
type GeminiExperimentHandler struct {
	model *genai.GenerativeModel
	ctx   context.Context
}

// NewGeminiExperimentHandler creates a new instance of GeminiExperimentHandler and returns it as an IExperimentHandler
func NewGeminiExperimentHandler() contracts.IExperimentHandler {
	return &GeminiExperimentHandler{
		model: GetTextModel(), // Assuming `model` is accessible from this scope
		ctx:   context.Background(),
	}
}

func (eh *GeminiExperimentHandler) GetExperimentsForChapter(generateDTO *dtos.GenerateExperimentDTO, chapterName string) ([]*domain.Experiment, error) {
	resp, err := eh.model.GenerateContent(eh.ctx, genai.Text("List possible experiments for the chapter "+ chapterName + ` return the response in json format [{"ExperimentName": "Experiment 1: experiment name"}, {"ExperimentName": "Experiment 2: experiment name"}, {"ExperimentName": "Experiment 3: experiment name"}]`))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Get the response content
	responseContent := resp.Candidates[0].Content.Parts[0]

	// Convert the response content to byte
	responseContentJSON, _ := json.Marshal(responseContent)

	// Convert the response content from byte to string
	var responseContentString string
	err = json.Unmarshal(responseContentJSON, &responseContentString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Convert the response content to a list of domain.Experiment
	var experiments []*domain.Experiment
	err = json.Unmarshal([]byte(responseContentString[8:len(responseContentString)-3]), &experiments)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Set the SubjectID and ChapterID for each experiment
	for _, exp := range experiments {
		exp.SubjectID = generateDTO.SubjectID
		exp.ChapterID = generateDTO.ChapterID
	}

	return experiments, nil
}


func (eh *GeminiExperimentHandler) GetExperimentContent(experimentName string, promptMessage string) ([]byte, error) {
	resp, err := eh.model.GenerateContent(eh.ctx, genai.Text("Generate content for the experiment " + experimentName + promptMessage))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Get the response content
	responseContent := resp.Candidates[0].Content.Parts[0]

	// Convert the response content to byte
	responseContentJSON, _ := json.Marshal(responseContent)

	// Convert the response content from byte to string
	var responseContentString string
	err = json.Unmarshal(responseContentJSON, &responseContentString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return []byte(responseContentString), nil
}

