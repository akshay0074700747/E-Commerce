package cronejobs

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UnblockUsers struct {
	DB *gorm.DB
}

func NewUnblockUsers(db *gorm.DB) *UnblockUsers {

	return &UnblockUsers{DB: db}

}

func (un *UnblockUsers) Start() {

	ticker := time.NewTicker(5 * time.Minute)

	go func() {
		for range ticker.C {

			if err := un.DB.Exec(`UPDATE users SET isblocked = false WHERE unblock_time < NOW();`).Error; err != nil {
				fmt.Println(err.Error())
			}

			if err := un.DB.Exec(`UPDATE admins SET isblocked = false WHERE unblock_time < NOW();`).Error; err != nil {
				fmt.Println(err.Error())
			}

			if err := un.DB.Exec(`DELETE FROM discounts WHERE end_date < NOW();`).Error; err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println("croneeee woooorked.........................................")

		}
	}()

}
