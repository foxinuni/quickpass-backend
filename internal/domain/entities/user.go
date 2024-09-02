package entities

type User struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Number string `json:"number"`
}

func NewUser(userID int, email string, number string) *User {
	return &User{
		UserID: userID,
		Email:  email,
		Number: number,
	}
}

func (u *User) GetUserID() int {
	return u.UserID
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetNumber() string {
	return u.Number
}
