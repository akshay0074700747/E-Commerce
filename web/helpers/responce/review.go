package responce

import "time"

type ReviewResponce struct {
	ID          uint      `json:"id"`
	Product     uint      `json:"product"`
	ReviewedBy  string    `json:"reviewed_by"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
}
