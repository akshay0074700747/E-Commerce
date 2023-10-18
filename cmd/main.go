package main

import (
	"ecommerce/web/config"
	"ecommerce/wireeeeeee"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		panic(err.Error())
	}

	server, err := wireeeeeee.InitializeAPI1(config)

	if err != nil {
		panic("Couldnt start server")
	}

	server.Start()

}
