package db

import (
	"github.com/nandobas/twitter/domain/message"
)

type MessageRepositoryMock struct {
	CreateFunc          func(m message.Message) error
	GetMessageByIDFunc  func(messageID string) (message.Message, error)
	LoadAllMessagesFunc func() ([]message.Message, error)
}

func (mr *MessageRepositoryMock) Create(m message.Message) error {
	if mr.CreateFunc != nil {
		return mr.CreateFunc(m)
	}
	return nil
}

func (mr *MessageRepositoryMock) GetMessageByID(messageID string) (message.Message, error) {
	if mr.GetMessageByIDFunc != nil {
		return mr.GetMessageByIDFunc(messageID)
	}
	return message.Message{}, nil
}

func (mr *MessageRepositoryMock) LoadAllMessages() ([]message.Message, error) {
	if mr.LoadAllMessagesFunc != nil {
		return mr.LoadAllMessagesFunc()
	}
	var list []message.Message
	return list, nil
}
