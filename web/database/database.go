package database

import (
	"ecommerce/internal/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Group_tables interface {
	Migrate_me()
}

// type DBConn struct {
// 	DB *gorm.DB
// }

func Connect_to(envs map[string]string) *gorm.DB {

	var db *gorm.DB
	var err error

	if db, err = gorm.Open(postgres.Open(envs["DATABASE_ADDR"]), &gorm.Config{}); err != nil {
		panic("cannot connect to the databse...")
	} else {
		Migrte_all(db,&entities.User{})
		return db
	}
}

func Migrte_all(dbconn *gorm.DB, models ...Group_tables) {
	for _, model := range models {
		dbconn.AutoMigrate(model)
	}
}
