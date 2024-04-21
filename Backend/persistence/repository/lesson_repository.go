package repositories

import (
	"os"
	"path/filepath"
	"strconv"

	"gorm.io/gorm"
	"hackathon.com/leariningApp/domain"
	// "log"
)

type LessonRepository struct {
	*GenericRepository[domain.Lesson]
}

func NewLessonRepository(GetDb func() *gorm.DB) *SubjectRepository {
	return &SubjectRepository{
		GenericRepository: NewGenericRepository[domain.Subject](GetDb),
	}
}

func (repository *LessonRepository) CreateLesson(name string, chapterID uint) (*domain.Lesson, error) {
	lesson := domain.Lesson{
		Name:      name,
		ChapterID: chapterID,
	}

	// Generate content for the lesson
	content := `# Introduction to the Nervous System

## Structure of the Nervous System

The nervous system is a complex network of cells that allows communication between different parts of the body. It can be divided into two main parts: the central nervous system (CNS) and the peripheral nervous system (PNS).

The CNS consists of the brain and spinal cord. The brain is the command center of the nervous system, responsible for processing information, controlling body movements, and regulating vital functions such as breathing and heartbeat. The spinal cord serves as a pathway for transmitting signals between the brain and the rest of the body.

The PNS includes all nerves outside the CNS and is further divided into the somatic nervous system and the autonomic nervous system. The somatic nervous system controls voluntary movements and receives sensory information from the external environment, while the autonomic nervous system regulates involuntary functions such as heart rate, digestion, and breathing.

![Image](https://example.com/image.jpg)

## Function of the Nervous System

The nervous system functions by transmitting electrical impulses, known as action potentials, between nerve cells called neurons. Neurons are the basic building blocks of the nervous system and are specialized for transmitting and processing information.

1. **Sensory Function**: Sensory neurons detect stimuli from the external environment or internal body conditions and transmit this information to the brain for processing. Examples of sensory stimuli include touch, temperature, pain, and sound.

2. **Motor Function**: Motor neurons carry signals from the brain and spinal cord to muscles and glands, controlling voluntary and involuntary movements. Motor neurons activate muscles to contract and produce movement, allowing us to walk, talk, and perform other activities.

3. **Integration and Processing**: The brain integrates sensory information, processes it, and generates appropriate responses. Complex networks of neurons form circuits and pathways that process information and control specific functions such as memory, emotion, and learning.

![Image](https://example.com/image.jpg)

## Visualization of the Nervous System

Below are visualizations of key components of the nervous system:

1. **Brain Anatomy**: ![Image](https://example.com/image.jpg)

2. **Spinal Cord Structure**: ![Image](https://example.com/image.jpg)

3. **Autonomic Nervous System**: ![Image](https://example.com/image.jpg)

## Introduction

Have you ever wondered how you can taste your favorite food, feel the sunshine on your skin, or jump out of the way of a sudden loud noise? It's all thanks to your incredible nervous system! This complex network acts as your body's control center, constantly gathering information from the world around you and sending messages to make your body react.

## Learning Objectives

- Identify the two main parts of the nervous system.
- Explain the function of the central nervous system and peripheral nervous system.
- Describe how nerve cells transmit messages throughout the body.

## Conclusion

The nervous system is a remarkable and intricate system that enables communication and coordination throughout the body. Understanding its structure and function is essential for comprehending how the body responds to various stimuli and maintains homeostasis. By learning about the nervous system, we gain insights into the complexities of human physiology and the interconnectedness of our bodily functions.`

	// Create a new Markdown file and write the content to it
	contentFile, err := os.Create(filepath.Join("lessons", strconv.FormatUint(uint64(lesson.ID), 10)+".md"))
	if err != nil {
		return nil, err

	}
	defer contentFile.Close()

	_, err = contentFile.WriteString(content)
	if err != nil {
		return nil, err
	}

	// Store the link to the Markdown file in the database
	lesson.ContentLink = contentFile.Name()

	response, err := repository.Create(&lesson)

	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
{
  "lessonTitle": "Introduction to the Nervous System",
  "lessonContent": [
    {
      "sectionTitle": "Structure of the Nervous System",
      "sectionContent": "The nervous system is a complex network of cells that allows communication between different parts of the body. It can be divided into two main parts: the central nervous system (CNS) and the peripheral nervous system (PNS).\n\nThe CNS consists of the brain and spinal cord. The brain is the command center of the nervous system, responsible for processing information, controlling body movements, and regulating vital functions such as breathing and heartbeat. The spinal cord serves as a pathway for transmitting signals between the brain and the rest of the body.\n\nThe PNS includes all nerves outside the CNS and is further divided into the somatic nervous system and the autonomic nervous system. The somatic nervous system controls voluntary movements and receives sensory information from the external environment, while the autonomic nervous system regulates involuntary functions such as heart rate, digestion, and breathing.",
      "images": [
		"https://www.google.com/url?sa=i&url=https%3A%2F%2Fvisualsonline.cancer.gov%2Fdetails.cfm%3Fimageid%3D9115&psig=AOvVaw1SFLk05Ryizx-g87OBkCJX&ust=1713711841855000&source=images&cd=vfe&opi=89978449&ved=0CBIQjRxqFwoTCNCPjp-I0YUDFQAAAAAdAAAAABAE",
		"https://www.google.com/imgres?q=**%5BImage%20of%20Motor%20Neurons%5D**%20(Find%20images%20on%20sites%20like%20Khan%20Academy%20or%20explorevisiblebody.com)%20%2C%20%2F%2F%20Replace%20with%20found%20image%20URL&imgurl=https%3A%2F%2Fc8.alamy.com%2Fcomp%2FT802HR%2Fmotor-neuron-and-muscle-fiber-illustration-T802HR.jpg&imgrefurl=https%3A%2F%2Fwww.alamy.com%2Fstock-photo%2Fmotor-neuron.html&docid=lE3wUBG1YrsWEM&tbnid=RYp0RrB8h18H7M&vet=12ahUKEwjMgsTPiNGFAxV9F1kFHVQpDQQQM3oECGYQAA..i&w=1300&h=1178&hcb=2&ved=2ahUKEwjMgsTPiNGFAxV9F1kFHVQpDQQQM3oECGYQAA"
      ]
    },
    {
      "sectionTitle": "Function of the Nervous System",
      "sectionContent": "The nervous system functions by transmitting electrical impulses, known as action potentials, between nerve cells called neurons. Neurons are the basic building blocks of the nervous system and are specialized for transmitting and processing information.\n\n1. **Sensory Function**: Sensory neurons detect stimuli from the external environment or internal body conditions and transmit this information to the brain for processing. Examples of sensory stimuli include touch, temperature, pain, and sound.\n\n2. **Motor Function**: Motor neurons carry signals from the brain and spinal cord to muscles and glands, controlling voluntary and involuntary movements. Motor neurons activate muscles to contract and produce movement, allowing us to walk, talk, and perform other activities.\n\n3. **Integration and Processing**: The brain integrates sensory information, processes it, and generates appropriate responses. Complex networks of neurons form circuits and pathways that process information and control specific functions such as memory, emotion, and learning.",
      "images": [
        "**[Image of Sensory Neurons]** (Find images on sites like McGraw-Hill Education or explorevisiblebody.com)", // Replace with found image URL
        "**[Image of Motor Neurons]** (Find images on sites like Khan Academy or explorevisiblebody.com)", // Replace with found image URL
        "**[Image of Brain Processing]** (Find images on sites like National Institutes of Health (.gov) websites or McGraw-Hill Education)" // Replace with found image URL
      ]
    },
    {
      "sectionTitle": "Visualization of the Nervous System",
      "sectionContent": "Below are visualizations of key components of the nervous system:\n\n1. **Brain Anatomy**: [Insert image/diagram of brain anatomy]\n\n2. **Spinal Cord Structure**: [Insert image/diagram of spinal cord structure]\n\n3. **Autonomic Nervous System**: [Insert image/diagram of autonomic nervous system]",
      "images": [
        "", // Replace with description
        "", // Replace with description
        "" // Replace with description
      ]
    },
    {
      "sectionTitle": "Introduction",
      "sectionContent": "Have you ever wondered how you can taste your favorite food, feel the sunshine on your skin, or jump out of the way of a sudden loud noise? It's all thanks to your incredible nervous system! This complex network acts as your body's control center, constantly gathering information from the world around you and sending messages to make your body react."
    },
    {
      "sectionTitle": "Learning Objectives",
      "sectionContent": [
        "Identify the two main parts of the nervous system.",
        "Explain the function of the central nervous system and peripheral nervous system.",
        "Describe how nerve cells transmit messages throughout the body."
      ]
    },
    {
      "sectionTitle": "Conclusion",
      "sectionContent": "The nervous system is a remarkable and intricate system that enables communication and coordination throughout the body. Understanding its
*/
