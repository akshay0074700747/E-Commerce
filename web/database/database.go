package database

import (
	"ecommerce/web/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Group_tables interface {
	Migrate_me()
}

type ItsaDatabase interface {
	Returnconn() *gorm.DB
}

type DBConn struct {
	DB *gorm.DB
}

func Connect_to() *DBConn {

	val, err := config.LoadEnv("DATABASE_ADDR")
	if err != nil {
		panic("Coldnt load env filess")
	}
	var db *gorm.DB

	if db, err = gorm.Open(postgres.Open(val["DATABASE_ADDR"]), &gorm.Config{}); err != nil {
		panic("cannot connect to the databse...")
	}

	fmt.Println("connected to the databse successfully ... ")

	return &DBConn{DB: db}
}

func (dbconn *DBConn) Migrte_all(models ...Group_tables) {
	for _, model := range models {
		dbconn.DB.AutoMigrate(model)
	}
}

func (dbconn *DBConn) Returnconn() *gorm.DB {
	return dbconn.DB
}
