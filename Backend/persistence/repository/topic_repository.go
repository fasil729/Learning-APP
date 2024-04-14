package repositories

import (
	"gorm.io/gorm"
	"hackathon.com/leariningApp/domain"
	"hackathon.com/leariningApp/persistence/migrations"
	// "log"
)

type TopicRepository struct {
	*GenericRepository[domain.Topic]
}

func NewTopicRepository(GetDb func() *gorm.DB) *TopicRepository {
	return &TopicRepository{
		GenericRepository: NewGenericRepository[domain.Topic](GetDb),
	}
}

func (repository *TopicRepository) CreateTopic(TopicName string, Description string, UserID uint) (*domain.Topic, error) {
	topic := domain.Topic{
		TopicName:   TopicName,
		Description: Description,
		UserID:      UserID,
	}

	createdTopic, err := repository.Create(&topic)

	if err != nil {
		return nil, err
	}

	return createdTopic, nil
}

func CreateTopic(TopicName string, Description string, UserID uint) (*domain.Topic, error) {

	topic := &domain.Topic{
		TopicName:   TopicName,
		Description: Description,
		UserID:      UserID,
	}

	response := migrations.GetDB().Create(topic)
	error := response.Error

	if error != nil {
		return nil, error
	}

	// experiments, err := GetExperimentsByTopicID(TopicName)

	// if err != nil {
	// 	log.Printf("Error getting experiments from API: %v", err)
	// 	return nil, err
	// }

	// notes, err := GetNotesByTopicID(TopicName)

	// if err != nil {
	// 	log.Printf("Error getting notes from API: %v", err)
	// 	return nil, err
	// }

	// quizzes, err := api.GetQuizzesFromAPI(TopicName)

	// if err != nil {
	// 	log.Printf("Error getting quizzes from API: %v", err)
	// 	return nil, err
	// }
	// for _, experiment := range experiments {
	// 	dbExperiment := &domain.Experiment{TopicID: topic.ID /* other fields */}
	// 	migrations.GetDB().Create(dbExperiment)
	// }

	// for _, note := range notes {
	//     dbNote := &domain.Note{TopicID: topic.ID, /* other fields */}
	//     migrations.GetDB().Create(dbNote)
	// }
	// for _, quiz := range quizzes {
	//     dbQuiz := &domain.Quiz{TopicID: topic.ID, /* other fields */}
	//     migrations.GetDB().Create(dbQuiz)
	// }

	return topic, nil

}

func findTopicsOFUser(userId uint) ([]*domain.Topic, error) {
	var topics []*domain.Topic

	response := migrations.GetDB().Where("UserID = ?", userId).Find(&topics)
	error := response.Error

	if error != nil {
		return nil, error
	}

	return topics, nil
}

func searchTopicsByName(query string) ([]*domain.Topic, error) {
	var searchedTopics []*domain.Topic

	response := migrations.GetDB().Where("TopicName LIKE ?", "%"+query+"%").Find(&searchedTopics)
	error := response.Error

	if error != nil {
		return nil, error
	}

	return searchedTopics, nil

}
