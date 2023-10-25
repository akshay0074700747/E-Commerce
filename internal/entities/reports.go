package entities

type Reports struct {
	ID   uint   `gorm:"primaryKey;unique;not null"`
	Email string 
	Description string
}

func (report *Reports) Migrate_me() {
}