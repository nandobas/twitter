package timeline_test

import (
	"github.com/nandobas/twitter/domain/message"
	"github.com/nandobas/twitter/domain/user"
	"github.com/nandobas/twitter/infra/db"
	"github.com/nandobas/twitter/usecase/timeline"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TimelineTestSuite struct {
	suite.Suite
}

func TestTimeline(t *testing.T) {
	suite.Run(t, new(TimelineTestSuite))
}

func (t *TimelineTestSuite) TestTimeline_WhenLoadTimeline_ExpectedMessageList() {
	var list []message.Message
	list = append(list, message.Message{
		ID:      "1",
		UserID:  "123456",
		Message: "Hello Word",
	})
	list = append(list, message.Message{
		ID:      "2",
		UserID:  "123456",
		Message: "Bye",
	})

	messageRepository := &db.MessageRepositoryMock{
		LoadAllMessagesFunc: func() ([]message.Message, error) {
			return list, nil
		},
	}
	userRepository := &db.UserRepositoryMock{
		GetUserByIDFunc: func(userID string) (user.User, error) {
			return user.User{
				ID:   "123456",
				Name: "Basilio",
			}, nil
		},
	}

	messageTimeline := timeline.MessageTimeline{
		MessageRepository: messageRepository,
		UserRepository:    userRepository,
	}
	pageInput := timeline.PageInput{
		Page:   0,
		UserID: "123456",
	}
	expectedTotalMessages := 2

	givenTimeline, err := messageTimeline.LoadTimeline(pageInput)
	t.NoError(err)

	t.Equal(givenTimeline.TotalMessages, expectedTotalMessages)

}
