package postgres

import "api/models"

func GetChatHistory() ([]models.Message, error) {
	var messages []models.Message
	err := Database.Select(&messages, "select * from messages order by id;")
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func GetRecent(upto uint64) ([]models.Message, error) {
	var messages []models.Message
	err := Database.Select(&messages, "select * from messages order by id desc limit $1;", int(upto))
	if err != nil {
		return nil, err
	}

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}
