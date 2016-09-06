package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	ProductID uint
	UserID    uint
	Number    uint
}

func CreateOrderTable() {
	db := getConnectionDB()
	defer db.Close()

	if !db.HasTable(&Order{}) {
		db.CreateTable(&Order{})
	}
}
