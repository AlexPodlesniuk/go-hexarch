package domain

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	ReqNum int    `json:"reqNum"`
}

func NewUser(id string, name string) *User {
	return &User{ID: id, Name: name, ReqNum: 0}
}

func (u *User) updateCounter() {
	u.ReqNum++
}
