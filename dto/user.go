package dto

type User struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}
