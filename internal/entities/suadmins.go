package entities

type SuperAdmins struct {
	ID       uint   `gorm:"primaryKey;unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

func (suadmin *SuperAdmins) Migrate_me() {
}
