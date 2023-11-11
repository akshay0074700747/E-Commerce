package main

import (
	"ecommerce/web/config"
	cronejobs "ecommerce/web/crone_jobs"
	"ecommerce/web/database"
	"ecommerce/wireeeeeee"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		panic(err.Error())
	}

	togglecrone := make(chan bool, 6)
	listencrone := make(chan int, 1)

	server, err := wireeeeeee.InitializeAPI(config, togglecrone, listencrone)

	if err != nil {
		panic("Couldnt start server")
	}

	cronedb := database.Connect_to(config)

	unblock_crone := cronejobs.NewUnblockUsers(cronedb)

	go unblock_crone.Start(togglecrone, listencrone)

	server.Start()

}
