package message

type Repository interface {
	Create(m Message) error
	GetMessageByID(messageID string) (Message, error)
	LoadAllMessages() ([]Message, error)
}
