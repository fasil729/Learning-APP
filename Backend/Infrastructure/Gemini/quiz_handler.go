package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	contracts "Brilliant/application/contracts/gemini"
	dtos "Brilliant/application/dtos/quiz"

	"github.com/google/generative-ai-go/genai"
)

// GeminiQuizHandler implements the IGeminiQuizHandler interface
type GeminiQuizHandler struct {
	model *genai.GenerativeModel
	ctx   context.Context
}

// NewGeminiQuizHandler creates a new instance of GeminiQuizHandler
func NewGeminiQuizHandler() contracts.IQuizHandler {
	return &GeminiQuizHandler{
		model: GetTextModel(), // Assuming `model` is accessible from this scope
		ctx:   context.Background(),
	}
}

func (qh *GeminiQuizHandler) GenerateQuiz(dto *dtos.GenerateQuizDTO) ([]dtos.Quiz, error) {
	// Generate a prompt based on the topics
	quizPrompt := "Generate " + strconv.Itoa(dto.NumberOfQuizzes) + " quiz question(s) for the topic(s) " + topicsList(dto.Topics) + ". " +
		"Each quiz should include a question, four options, and an explanation for the correct answer. " +
		"Among the options, ensure that only one is the correct answer, marked with 'is_answer: true', while all other options are marked 'is_answer: false'. " +
		"Return the response in the following JSON format: " +
		`[
    {
        "question": "Insert the quiz question here",
        "options": [
            {"text": "Option 1", "is_answer": false},
            {"text": "Option 2", "is_answer": false},
            {"text": "Option 3", "is_answer": true}, // This is the correct answer
            {"text": "Option 4", "is_answer": false}
        ],
        "explanation": "Provide a detailed explanation for why the correct answer is correct"
    },
    {
        // More quizzes can be added here
    }
].`

	resp, err := qh.model.GenerateContent(qh.ctx, genai.Text(quizPrompt))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// debugging to see what will be returned from gemini
	// Number_of_candidates := len(resp.Candidates)

	// fmt.Println("Number of candidates: ", Number_of_candidates)
	// for i := 0; i < Number_of_candidates; i++ {
	// 	fmt.Println("candidate", resp.Candidates[i].Content)

	// }

	quizzes := make([]dtos.Quiz, dto.NumberOfQuizzes)

	responseContent := resp.Candidates[0].Content.Parts[0]
	responseContentJSON, _ := json.Marshal(responseContent)

	var responseContentString string
	err = json.Unmarshal(responseContentJSON, &responseContentString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// debuging to see what will be returned from gemini
	fmt.Println("responseContentString: ", responseContentString)

	err = json.Unmarshal([]byte(responseContentString), &quizzes)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// debug

	return quizzes, nil
}

func topicsList(topics []string) string {
	result := ""
	for i, topic := range topics {
		if i > 0 {
			result += ", "
		}
		result += topic
	}
	return result
}
