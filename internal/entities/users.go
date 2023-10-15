package entities

type User struct {
	Id       string `gorm:"primaryKey"`
	Email    string
	Password string
	Mobile   string
	Name     string
}

func (user *User) Migrate_me() {
}