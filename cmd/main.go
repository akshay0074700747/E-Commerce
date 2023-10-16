package main

import (
	"ecommerce/web/config"
	"ecommerce/wireeeeeee"
)

func main() {

	envs, err := config.LoadEnv("DATABASE_ADDR", "EMAIL", "PASSWORD", "SECRET")

	if err != nil {
		panic("Cannot connect to the database...")
	}

	server,err := wireeeeeee.InitializeAPI(envs)

	if err != nil {
		panic("Couldnt start server")
	}

	server.Start()

}
