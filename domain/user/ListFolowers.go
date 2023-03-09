package user

import "fmt"

type ListFollowing struct {
	UserID string
	List   []string
}

func (lf ListFollowing) AddUser(userID string) error {
	find := lf.findUserInFollowing(userID)
	if find {
		return fmt.Errorf("user already available in following %s", userID)
	}
	return nil
}

func (lf ListFollowing) findUserInFollowing(userID string) bool {
	for _, findUserID := range lf.List {
		if findUserID == userID {
			return true
		}
	}
	return false
}
