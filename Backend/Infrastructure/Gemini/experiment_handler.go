package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	contracts "Brilliant/application/contracts/gemini"
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

func (eh *GeminiExperimentHandler) GetExperimentsForChapter(generateDTO *dtos.GenerateExperimentDTO) ([]domain.Experiment, error) {
	resp, err := eh.model.GenerateContent(eh.ctx, genai.Text("List possible experiments for the chapter "+generateDTO.ChapterName+` return the response in json format [{"ExperimentName": "Experiment 1: experiment name"}, {"ExperimentName": "Experiment 2: experiment name"}, {"ExperimentName": "Experiment 3: experiment name"}] make them specific for the chapter only`))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Get the response content
	responseContent := resp.Candidates[0].Content.Parts[0]

	// Convert the response content to JSON
	responseContentJSON, _ := json.Marshal(responseContent)

	// Convert the response content from JSON to string
	var responseContentString string
	err = json.Unmarshal(responseContentJSON, &responseContentString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// debug
	fmt.Println("responseContentString: ", responseContentString)

	// Parse the JSON response directly into a slice of structs
	var experimentsJSON []struct {
		ExperimentName string `json:"ExperimentName"`
	}
	if err := json.Unmarshal([]byte(responseContentString[8:len(responseContentString)-3]), &experimentsJSON); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// debug
	fmt.Println("experimentsJSON: ", experimentsJSON)

	// Build the domain.Experiment objects using the experiment names and SubjectID and ChapterID from generateDTO
	var experiments []domain.Experiment
	for _, exp := range experimentsJSON {
		experiments = append(experiments, domain.Experiment{
			ExperimentName: exp.ExperimentName,
			// SubjectID:      generateDTO.SubjectID,
			// ChapterID:      generateDTO.ChapterID,
		})
	}

	// debug
	fmt.Println("experiments: ", experiments)

	return experiments, nil
}

func (eh *GeminiExperimentHandler) GetExperimentContent(experimentName string, promptMessage string) ([]byte, error) {
	prompt := "Generate content for the experiment '" + experimentName + "'. " +
		"The following details should be included: " + promptMessage + ". " +
		"Use the following markdown format for content creation, incorporating the given structure: " +
		"\n\n# Experiment Title\n\n" +
		"## Introduction\n" +
		"Briefly describe the context and relevance of the experiment. What is it about? Why is it important?\n\n" +
		"## Objective\n" +
		"State the goal of the experiment. What do you aim to achieve or demonstrate?\n\n" +
		"## Materials\n" +
		"List the materials required for the experiment:\n" +
		"- Material 1\n" +
		"  <ImagePlaceholder Description=\"/path/to/material1.jpg\" width=\"718\" height=\"404\" alt=\"Description of Material 1\" />\n" +
		"- Material 2\n" +
		"  <ImagePlaceholder Description=\"/path/to/material2.jpg\" width=\"718\" height=\"404\" alt=\"Description of Material 2\" />\n\n" +
		"## Procedure\n" +
		"Describe the steps in the experiment:\n" +
		"1. Step 1: Describe the first step in detail.\n" +
		"   <ImagePlaceholder Description=\"/path/to/step1.jpg\" width=\"718\" height=\"404\" alt=\"Image of Step 1\" />\n" +
		"2. Step 2: Describe the second step in detail.\n" +
		"   <ImagePlaceholder Description=\"/path/to/step2.jpg\" width=\"718\" height=\"404\" alt=\"Image of Step 2\" />\n\n" +
		"## Safety Precautions\n" +
		"List any safety precautions or warnings to consider:\n" +
		"- Safety Tip 1\n" +
		"- Safety Tip 2\n\n" +
		"## Results\n" +
		"Describe the expected results of the experiment. What should the final outcome or data indicate?\n\n" +
		"## Discussion\n" +
		"Discuss the results and their significance. What do they mean in the context of the experiment's objective?\n\n" +
		"## Conclusion\n" +
		"Summarize the key findings and their broader implications. How do the results contribute to the understanding of the topic?\n\n" +
		"Ensure the content is cohesive, informative, and visually appealing. Include appropriate image placeholders in the specified format, and maintain logical transitions between sections."

	resp, err := eh.model.GenerateContent(eh.ctx, genai.Text(prompt))
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
