package main

import (
	"ecommerce/internal/entities"
	"ecommerce/web/config"
	"ecommerce/web/database"
)

func main() {

	envs, err := config.LoadEnv("DATABASE_ADDR")

	if err != nil {
		panic("Cannot connect to the database...")
	}

	addr := envs["DATABASE_ADDR"]

	database.Connect_to(addr)

	database.Migrte_all(&entities.User{})

}