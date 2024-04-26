package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"Brilliant/application/contracts/gemini"

	"github.com/google/generative-ai-go/genai"
)

// GeminisubjectHandler implements ISubjectHandler
type GeminiSubjectHandler struct {
	model *genai.GenerativeModel
	ctx  context.Context
}

// NewGeminiSubjectHandler creates a new instance of GeminisubjectHandler and returns it as an IGeminiSubjectHandler
func NewGeminiSubjectHandler() contracts.IGeminiSubjectHandler {
	return &GeminiSubjectHandler{
		model: GetTextModel(), // Assuming `model` is accessible from this scope
		ctx:  context.Background(),
	}
}

func (sh *GeminiSubjectHandler) GenerateTableOfContent(subjectName string) ([]contracts.Chapter, error) {
    resp, err := sh.model.GenerateContent(sh.ctx, genai.Text("Generate a table of content for the subject "+subjectName))
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    // Debugging to see what will be returned from Gemini
    numberOfCandidates := len(resp.Candidates)
    fmt.Println("Number of candidates:", numberOfCandidates)
    for i := 0; i < numberOfCandidates; i++ {
        numberOfParts := len(resp.Candidates[0].Content.Parts)
        for j := 0; j < numberOfParts; j++ {
            fmt.Println("candidate", resp.Candidates[0].Content.Parts[j])
        }
    }

    responseContent := resp.Candidates[0].Content.Parts[0]

    var chapterString string
    var chapters []contracts.Chapter

    responseContentJSON, _ := json.Marshal(responseContent)
    fmt.Println("responseContentJSON:", responseContentJSON)

    err = json.Unmarshal([]byte(responseContentJSON), &chapterString)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    fmt.Println("chapterString):", chapterString)

    err = json.Unmarshal([]byte(chapterString[8:len(chapterString)-3]), &chapters)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    fmt.Println("chapters:", chapters)

    return chapters, nil
}


func (sh *GeminiSubjectHandler) GenerateLessonDetailContent(lessonName string) ([]byte, error) {
	resp, err := sh.model.GenerateContent(sh.ctx, genai.Text("Generate detailed content for the lesson "+lessonName))
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


