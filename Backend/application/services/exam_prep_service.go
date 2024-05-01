package services

import (
	contract "Brilliant/application/contracts/gemini"
	dtos "Brilliant/application/dtos/exam_prep"

	// "fmt"

	"github.com/pkg/errors"
)

type ExamPrepService struct {
	GeminiExamPrepHandler contract.IExamPrepHandler
}

func NewExamPrepService(geminiExamPrepHandler contract.IExamPrepHandler) *ExamPrepService {
	return &ExamPrepService{
		GeminiExamPrepHandler: geminiExamPrepHandler,
	}
}

func (eps *ExamPrepService) GenerateExamPrepMaterial(dto *dtos.GenerateExamPrepDTO) ([]byte, error) {
	if len(dto.Topics) == 0 {
		return []byte{}, errors.New("topics list is empty")
	}

	// Generate exam preparation material using Gemini
	candidates, err := eps.GeminiExamPrepHandler.GenerateExamPrep(dto)

	if err != nil {
		return []byte{}, errors.Wrap(err, "failed to generate exam preparation material")
	}

	// Select the best candidate

	// Sample content generation
	// content := fmt.Sprintf("# Exam Preparation Material\n\n**Prompt:** %s\n\n**Read Time:** %d\n\n", dto.Prompt, dto.ReadTime)
	// for _, topic := range dto.Topics {
	// 	content += fmt.Sprintf("## %s\n\n", topic)
	// 	content += "### Sample Explanation:\n\n"
	// 	content += "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed et nisi turpis. Integer malesuada semper justo vel tempus.\n\n"

	// 	content += "### Sample Quiz:\n\n"
	// 	content += "#### Question:\n\n"
	// 	content += "What is the main function of the circulatory system?\n\n"

	// 	content += "#### Options:\n\n"
	// 	content += "- Option 1: To digest food\n"
	// 	content += "- Option 2: To exchange gases\n"
	// 	content += "- Option 3: To pump blood\n"
	// 	content += "- Option 4: To produce energy\n\n"

	// 	content += "#### Explanation:\n\n"
	// 	content += "The correct answer is Option 3: To pump blood. The circulatory system is responsible for transporting nutrients, gases, and waste products throughout the body. The heart pumps blood through a network of blood vessels to deliver these substances to cells and tissues.\n\n"
	// }

	// Convert to Markdown format
	return candidates, nil
}
