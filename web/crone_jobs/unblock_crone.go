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

func (un *UnblockUsers) Start(togglecrone chan bool, listencrone chan int) {

	ticker := time.NewTicker(5 * time.Minute)

	for {
		<-listencrone
		un.Crone(ticker, togglecrone)
	}

}

func (un *UnblockUsers) Crone(ticker *time.Ticker, togglecrone chan bool) {

	fmt.Println("jasdfTYDHKJASDVHVADGluidfjh")

	for range ticker.C {

		select {
		case val := <-togglecrone:
			if !val {
				fmt.Println(val)
				return
			}
		default:
			fmt.Println("continue...")
		}

		if err := un.DB.Exec(`UPDATE users SET isblocked = false WHERE unblock_time < NOW();`).Error; err != nil {
			fmt.Println(err.Error())
		}

		if err := un.DB.Exec(`UPDATE admins SET isblocked = false WHERE unblock_time < NOW();`).Error; err != nil {
			fmt.Println(err.Error())
		}

		if err := un.DB.Exec(`DELETE FROM discounts WHERE end_date < NOW();`).Error; err != nil {
			fmt.Println(err.Error())
		}

		if err := un.DB.Exec(`UPDATE orders SET is_shipped = true WHERE shipment_date < NOW() AND is_cancelled = false AND return_status = false AND (cod = true OR recieved_payment = true)`).Error; err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("croneeee woooorked.........................................")

	}
}
