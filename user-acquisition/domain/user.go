package domain

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUser(id, name string) *User {
	return &User{id, name}
}
