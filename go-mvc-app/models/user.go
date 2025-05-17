package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func GetAllUsers() []User {
	return users
}

func AddUser(user User) {
	users = append(users, user)
}
