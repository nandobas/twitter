package timeline

import (
	"fmt"
	"github.com/nandobas/twitter/domain/message"
	"github.com/nandobas/twitter/domain/user"
	"time"
)

type MessageDto struct {
	ID       string
	UserID   string
	UserName string
	Message  string
}

type MessagesOutput struct {
	TotalMessages int
	Page          int
	NextPage      int
	LastUpdate    time.Time
	ListMessages  []MessageDto
}

type PageInput struct {
	Page   int
	UserID string
}

type MessageTimeline struct {
	MessageRepository message.Repository
	UserRepository    user.Repository
}

func (mt MessageTimeline) LoadTimeline(pageInput PageInput) (MessagesOutput, error) {
	messages, err := mt.MessageRepository.LoadAllMessages()
	var messageOutput MessagesOutput
	if err != nil {
		return MessagesOutput{}, fmt.Errorf("cannot load messages: %w", err)
	}
	messageOutput.Page = pageInput.Page
	messageOutput.TotalMessages = len(messages)
	messageOutput.NextPage = 0
	messageOutput.LastUpdate = time.Now()

	for _, givenMessage := range messages {
		messageDto, err := mt.newMessageDto(givenMessage)
		if err != nil {
			return MessagesOutput{}, fmt.Errorf("cannot get user in newMessageDto: %w", err)
		}
		messageOutput.ListMessages = append(messageOutput.ListMessages, messageDto)
	}
	return messageOutput, nil
}

func (mt MessageTimeline) newMessageDto(givenMessage message.Message) (MessageDto, error) {
	user, err := mt.UserRepository.GetUserByID(givenMessage.UserID)
	if err != nil {
		return MessageDto{}, fmt.Errorf("cannot get user in newMessageDto: %w", err)
	}
	return MessageDto{
		ID:       givenMessage.ID,
		UserID:   givenMessage.UserID,
		UserName: user.Name,
		Message:  givenMessage.Message,
	}, nil
}
