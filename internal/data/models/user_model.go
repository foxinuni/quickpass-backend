package models

type User struct {
	UserID int
	Email  string
	Number string
}

func NewUser(userID int, email string, number string) *User {
	return &User{
		UserID: userID,
		Email:  email,
		Number: number,
	}
}
