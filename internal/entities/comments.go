package entities

import "time"

type Comments struct {
	ID uint `gorm:"primaryKey;unique;not null"`
	Review Review `gorm:"foreignKey:review_id"`
	CommentedBy string `gorm:"not null"`
	CommentDesc string
	CreatedAt time.Time
}

func (comment *Comments) Migrate_me() {
}