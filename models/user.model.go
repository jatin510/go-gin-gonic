package models

type User struct {
	ID       string `json:"user_id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) GetByID(id string) (*User, error) {
	var user *User
	var err error

	return user, err
}
