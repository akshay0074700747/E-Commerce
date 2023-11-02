package database

import (
	"ecommerce/internal/entities"
	"ecommerce/web/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Group_tables interface {
	Migrate_me()
}

func Connect_to(config config.Config) *gorm.DB {

	var db *gorm.DB
	var err error

	fmt.Println(config.DATABASE_ADDR)

	if db, err = gorm.Open(postgres.Open(config.DATABASE_ADDR), &gorm.Config{}); err != nil {
		panic(err.Error())
	} else {
		Migrte_all(db,
			&entities.User{},
			&entities.Admins{},
			&entities.SuperAdmins{},
			&entities.Brands{},
			&entities.Categories{},
			// &entities.Products{},
			&entities.Comments{},
			&entities.Discount{},
			&entities.Review{},
			&entities.Reports{},
			&entities.Carts{},
			&entities.CartItems{},
			&entities.WishLists{},
			&entities.WishListItems{},
			&entities.Address{},
			&entities.Orders{},
			&entities.OrderItems{})
		return db
	}
}

func Migrte_all(dbconn *gorm.DB, models ...Group_tables) {
	for _, model := range models {
		dbconn.AutoMigrate(model)
	}
}
