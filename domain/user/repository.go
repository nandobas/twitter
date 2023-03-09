package user

type Repository interface {
	Create(u User) error
	GetUserByID(userID string) (User, error)
}
