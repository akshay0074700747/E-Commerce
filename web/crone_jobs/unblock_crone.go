package cronejobs

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type UnblockUsers struct {
	DB *gorm.DB
}

func NewUnblockUsers(db *gorm.DB) *UnblockUsers {

	return &UnblockUsers{DB: db}

}

func (un *UnblockUsers) Start() {

	c := cron.New()

	_, err := c.AddFunc("*/5 * * * *", func() {

		if err := un.DB.Exec(`UPDATE users SET isblocked = false WHERE unblock_time < NOW();`).Error; err != nil {
			fmt.Println(err.Error())
		}

		if err := un.DB.Exec(`UPDATE admins SET isblocked = false WHERE unblock_time < NOW();`).Error; err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("croneeee woooorked.........................................")

	})

	if err != nil {
		fmt.Println(err.Error())
	}

	c.Start()

}
