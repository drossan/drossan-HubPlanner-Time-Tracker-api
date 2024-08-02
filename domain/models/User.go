package models

type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"userName"`
	Password string `json:"password"`
}
