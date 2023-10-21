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

func (un *UnblockUsers) Start()  {
	
	c := cron.New()

	_,err := c.AddFunc("*/5 * * * * *",func() {

		un.DB.Raw(`UPDATE users SET isblocked = false WHERE unblock_time > NOW();`)

	})

	if err != nil {
		fmt.Println(err.Error())
	}

	c.Start()

}