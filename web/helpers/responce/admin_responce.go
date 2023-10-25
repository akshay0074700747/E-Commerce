package responce

import "time"

type AdminData struct {
	Id        uint   `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Isblocked bool   `json:"isblocked"`
	Reports   int    `json:"reports"`
}

type AdminsideUsersData struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	Isblocked bool      `json:"isblocked"`
	Reports   int       `json:"reports"`
	CreatedAt time.Time `json:"createdat"`
}
