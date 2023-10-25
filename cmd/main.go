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

	server, err := wireeeeeee.InitializeAPI(config)

	if err != nil {
		panic("Couldnt start server")
	}

	cronedb := database.Connect_to(config)

	unblock_crone := cronejobs.NewUnblockUsers(cronedb)

	unblock_crone.Start()

	server.Start()

}
