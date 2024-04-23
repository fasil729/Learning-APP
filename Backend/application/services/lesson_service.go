package services

import (
	contracts "Brilliant/application/contracts/persistence"
	dtos "Brilliant/application/dtos/lesson"
	"Brilliant/domain"
	"os"
)

type LessonService struct {
	LessonRepository contracts.ILessonRepository
}

func NewLessonService(lessonRepository contracts.ILessonRepository) *LessonService {
	return &LessonService{
		LessonRepository: lessonRepository,
	}
}

func (service *LessonService) CreateLesson(lessonDTO *dtos.LessonDTO) (*domain.Lesson, error) {
	lesson := &domain.Lesson{
		LessonName: lessonDTO.LessonName,
		ChapterID:  lessonDTO.ChapterID,
	}

	lesson, err := service.LessonRepository.Create(lesson)
	if err != nil {
		return nil, err
	}
	return lesson, nil

}

func (service *LessonService) UpdateLesson(lessonId uint, lessonDTO *dtos.LessonDTO) (*domain.Lesson, error) {
	lesson, err := service.LessonRepository.GetById(lessonId)
	if err != nil {
		return nil, err
	}

	lesson.LessonName = lessonDTO.LessonName

	lesson, err = service.LessonRepository.Update(lesson)
	if err != nil {
		return nil, err
	}
	return lesson, nil
}

func (service *LessonService) DeleteLesson(lessonId uint) error {
	lesson, err := service.LessonRepository.GetById(lessonId)
	if err != nil {
		return err
	}

	_, err = service.LessonRepository.Delete(lesson)
	if err != nil {
		return err
	}
	return nil
}

func (service *LessonService) GetLessonContent(LessonID uint) ([]byte, error) {
	lesson, err := service.LessonRepository.GetById(LessonID)
	if err != nil {
		return nil, err
	}

	if lesson.ContentLink == "" {
		// Gemini API call will be done here

		// For now, return a dummy experiment
		lesson.ContentLink = "data/experiment/generatedLesson.md"

		// Update the experiment with the content link
		if _, err := service.LessonRepository.Update(lesson); err != nil {
			return nil, err
		}
	}

	// Read the content of the JSON file
	content, err := os.ReadFile(lesson.ContentLink)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (service *LessonService) GenerateLessons(chapterID uint) ([]*domain.Lesson, error) {
	lessonNames := []string{
		"Introduction to the Nervous System",
		"Structure of the Nervous System",
		"Function of the Nervous System",
		"Visualization of the Nervous System",
		"Introduction",
		"Learning Objectives",
		"Conclusion",
	}

	var lessons []*domain.Lesson
	for _, lessonName := range lessonNames {
		lesson := domain.Lesson{
			LessonName: lessonName,
			ChapterID:  chapterID,
		}

		createdLesson, err := service.LessonRepository.Create(&lesson)
		if err != nil {
			return nil, err
		}

		lessons = append(lessons, createdLesson)
	}

	return lessons, nil
}

// 	// Generate content for the lesson
// 	content := `# Introduction to the Nervous System

// ## Structure of the Nervous System

// The nervous system is a complex network of cells that allows communication between different parts of the body. It can be divided into two main parts: the central nervous system (CNS) and the peripheral nervous system (PNS).

// The CNS consists of the brain and spinal cord. The brain is the command center of the nervous system, responsible for processing information, controlling body movements, and regulating vital functions such as breathing and heartbeat. The spinal cord serves as a pathway for transmitting signals between the brain and the rest of the body.

// The PNS includes all nerves outside the CNS and is further divided into the somatic nervous system and the autonomic nervous system. The somatic nervous system controls voluntary movements and receives sensory information from the external environment, while the autonomic nervous system regulates involuntary functions such as heart rate, digestion, and breathing.

// ![Image](https://example.com/image.jpg)

// ## Function of the Nervous System

// The nervous system functions by transmitting electrical impulses, known as action potentials, between nerve cells called neurons. Neurons are the basic building blocks of the nervous system and are specialized for transmitting and processing information.

// 1. **Sensory Function**: Sensory neurons detect stimuli from the external environment or internal body conditions and transmit this information to the brain for processing. Examples of sensory stimuli include touch, temperature, pain, and sound.

// 2. **Motor Function**: Motor neurons carry signals from the brain and spinal cord to muscles and glands, controlling voluntary and involuntary movements. Motor neurons activate muscles to contract and produce movement, allowing us to walk, talk, and perform other activities.

// 3. **Integration and Processing**: The brain integrates sensory information, processes it, and generates appropriate responses. Complex networks of neurons form circuits and pathways that process information and control specific functions such as memory, emotion, and learning.

// ![Image](https://example.com/image.jpg)

// ## Visualization of the Nervous System

// Below are visualizations of key components of the nervous system:

// 1. **Brain Anatomy**: ![Image](https://example.com/image.jpg)

// 2. **Spinal Cord Structure**: ![Image](https://example.com/image.jpg)

// 3. **Autonomic Nervous System**: ![Image](https://example.com/image.jpg)

// ## Introduction

// Have you ever wondered how you can taste your favorite food, feel the sunshine on your skin, or jump out of the way of a sudden loud noise? It's all thanks to your incredible nervous system! This complex network acts as your body's control center, constantly gathering information from the world around you and sending messages to make your body react.

// ## Learning Objectives

// - Identify the two main parts of the nervous system.
// - Explain the function of the central nervous system and peripheral nervous system.
// - Describe how nerve cells transmit messages throughout the body.

// ## Conclusion

// The nervous system is a remarkable and intricate system that enables communication and coordination throughout the body. Understanding its structure and function is essential for comprehending how the body responds to various stimuli and maintains homeostasis. By learning about the nervous system, we gain insights into the complexities of human physiology and the interconnectedness of our bodily functions.`

// 	// Create a new Markdown file and write the content to it
// 	contentFile, err := os.Create(filepath.Join("lessons", strconv.FormatUint(uint64(lesson.ID), 10)+".md"))
// 	if err != nil {
// 		return nil, err

// 	}
// 	defer contentFile.Close()

// 	_, err = contentFile.WriteString(content)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Store the link to the Markdown file in the database
// 	lesson.ContentLink = contentFile.Name()

// 	response, err := repository.Create(&lesson)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return response, nil
