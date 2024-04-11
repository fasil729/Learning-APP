package topicRepositories

import (
	"hackathon.com/leariningApp/domain"
	"hackathon.com/leariningApp/persistence/migrations"
)

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

	} else {
		return topic, nil
	}

}
