package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"Brilliant/application/contracts/gemini"
	"Brilliant/application/dtos/quiz"

	"github.com/google/generative-ai-go/genai"
)

// GeminiQuizHandler implements the IGeminiQuizHandler interface
type GeminiQuizHandler struct {
	model *genai.GenerativeModel
	ctx   context.Context
}

// NewGeminiQuizHandler creates a new instance of GeminiQuizHandler
func NewGeminiQuizHandler() contracts.IQuizHandler  {
	return &GeminiQuizHandler{
		model: GetTextModel(), // Assuming `model` is accessible from this scope
		ctx:   context.Background(),
	}
}

func (qh *GeminiQuizHandler) GenerateQuiz(dto *dtos.GenerateQuizDTO) ([]dtos.Quiz, error) {
	// Generate a prompt based on the topics
	quizPrompt := "Generate " + strconv.Itoa(dto.NumberOfQuizzes) + " quiz(es) for the topic(s) " + topicsList(dto.Topics) + ` return the response in json format [{
		Question:   "question",
		Options:    [{Text: "Option 1", IsAnswer: false}, {Text: "Option 2", IsAnswer: false}, {Text: "Option 3", IsAnswer: true}, {Text: "Option 4", IsAnswer: false}}],
		Explanation: "Explanation for the correct answer",
	}]`

	resp, err := qh.model.GenerateContent(qh.ctx, genai.Text(quizPrompt))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// debugging to see what will be returned from gemini
	Number_of_candidates := len(resp.Candidates)

	fmt.Println("Number of candidates: ", Number_of_candidates)
	for i := 0; i < Number_of_candidates; i++ {
		fmt.Println("candidate", resp.Candidates[i].Content)

	}

    
	quizzes := make([]dtos.Quiz, dto.NumberOfQuizzes)

	responseContent := resp.Candidates[0].Content.Parts[0]
	responseContentJSON, _ := json.Marshal(responseContent)

	var responseContentString  string
	err = json.Unmarshal(responseContentJSON, &responseContentString)
	if err != nil {	
		log.Fatal(err)
		return nil, err
	}

	// debuging to see what will be returned from gemini
	fmt.Println("responseContentString: ", responseContentString)

	err = json.Unmarshal([]byte(responseContentString[8:len(responseContentString)-3]), &quizzes)

	if err != nil {	
		log.Fatal(err)
		return nil, err
	}

	// debug 
	fmt.Println("quizzes: ", quizzes)


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
