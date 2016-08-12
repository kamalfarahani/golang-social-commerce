package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func init() {
	CreateTables()
}

func CreateTables() {
	CreateProductTable()
	CreateCategoryTable()
	CreateUserTable()
}

func getConnectionDB() *gorm.DB {
	db, err := gorm.Open(
		"mysql", "iris:iris#max@/iris_db?charset=utf8&parseTime=True&loc=Local")
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
