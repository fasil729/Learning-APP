package gemini

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	contracts "Brilliant/application/contracts/gemini"
	dtos "Brilliant/application/dtos/exam_prep"

	"github.com/google/generative-ai-go/genai"
)

// GeminiExamPrepHandler implements IExamPrepHandler
type GeminiExamPrepHandler struct {
	model *genai.GenerativeModel
	ctx   context.Context
}

// NewGeminiExamPrepHandler creates a new instance of GeminiExamPrepHandler and returns it as an IExamPrepHandler
func NewGeminiExamPrepHandler() contracts.IExamPrepHandler {
	return &GeminiExamPrepHandler{
		model: GetTextModel(), // Assuming `model` is accessible from this scope
		ctx:   context.Background(),
	}
}

func (eh *GeminiExamPrepHandler) GenerateExamPrep(dto *dtos.GenerateExamPrepDTO) ([]byte, error) {
	// Generate a prompt based on the topics and study time
	topics := strings.Join(dto.Topics, ", ")
	prompt := "Generate exam preparation markdown material based on the following topics:" + topics +
		"The material should cover key concepts and provide a concise summary of each topic. " +
		"Include sample explanations, example problems, study tips, and suggested resources for further reading. " +
		"Here are additional details: " +
		"- **Prompt**: " + dto.Prompt + ". " +
		"- **Expected Study Time**: " + strconv.Itoa(dto.ReadTime) + " minutes. " +
		"- **Markdown Format**: Please structure the content in markdown format with headings, bullet points, and code snippets for clarity. " +
		"Suggested structure: " +
		"1. **Introduction**: A brief overview of the exam preparation material. Tips on how to best use the material and manage study time. " +
		"2. **Exam Topics**: For each topic, provide a summary of key concepts, a concise explanation, example problems, and study tips. " +
		"   Example format for each topic: " +
		"   ```markdown\n" +
		"   ## {{topic_name}}\n\n" +
		"   ### Summary\n" +
		"   - Key concepts, definitions, and essential information.\n\n" +
		"   ### Explanation\n" +
		"   Provide a detailed explanation of the topic with practical examples.\n\n" +
		"   ### Sample Quiz\n" +
		"   #### Question:\n" +
		"   Write a question that assesses the topic.\n\n" +
		"   #### Options:\n" +
		"   - Option 1: Sample answer 1\n" +
		"   - Option 2: Sample answer 2\n" +
		"   - Option 3: Sample answer 3\n" +
		"   - Option 4: Sample answer 4\n\n" +
		"   #### Correct Answer and Explanation:\n" +
		"   Explain why the correct answer is valid.\n\n" +
		"   ### Study Tips\n" +
		"   Provide helpful study tips for mastering the topic.\n\n" +
		"   ```\n" +
		"3. **Study Schedule**: Suggest an efficient study schedule to divide the study time across the topics and offer advice on managing study sessions. " +
		"4. **Additional Resources**: Recommend resources for further reading and exploration of the exam topics. " +
		"Ensure the generated content follows this structure and provides clear, concise, and useful exam preparation material."

	resp, err := eh.model.GenerateContent(eh.ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// var responseStrings []string
	// for _, candidate := range resp.Candidates {
	// 	responseContentJSON, _ := json.Marshal(candidate.Content.Parts[0])

	// 	var responseContentString string
	// 	err = json.Unmarshal(responseContentJSON, &responseContentString)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return nil, err
	// 	}
	// 	responseStrings = append(responseStrings, responseContentString)
	// }

	// return responseStrings, nil

	responseContent := resp.Candidates[0].Content.Parts[0]
	responseContentJSON, _ := json.Marshal(responseContent)
	var responseContentString string
	err = json.Unmarshal(responseContentJSON, &responseContentString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return []byte(responseContentString), nil
}
