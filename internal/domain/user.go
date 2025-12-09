package domain

type User struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
