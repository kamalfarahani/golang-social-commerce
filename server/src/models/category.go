package models

import (
	"github.com/jinzhu/gorm"

	"../constants"
)

type Category struct {
	gorm.Model
	ImgUrl   string
	Name     string    `sql:"size:50;unique;unique_index"`
	Products []Product `gorm:"ForeignKey:CatID"`
}

func CreateCategoryTable() {
	db := getConnectionDB()
	defer db.Close()

	if !db.HasTable(&Category{}) {
		db.CreateTable(&Category{})
	}
}

func SaveCategory(name string, products []Product) {
	db := getConnectionDB()
	defer db.Close()
	db.Save(
		&Category{
			Name:     name,
			Products: products,
		})
}

func GetAllCategories() []Category {
	var catArr []Category

	db := getConnectionDB()
	defer db.Close()
	db.Find(&catArr)

	return catArr
}

func GetCategoryByName(name string) (*Category, error) {
	result := new(Category)

	db := getConnectionDB()
	defer db.Close()
	db.Where("name = ?", name).First(result)
	db.Model(result).
		Related(&result.Products, "products")

	if result.Name != "" {
		return result, nil
	}
	return result, constants.CAT_NAME_ERR
}

func GetCategoryByID(id uint) (*Category, error) {
	result := new(Category)

	db := getConnectionDB()
	defer db.Close()
	db.First(result, id)
	db.Model(result).
		Related(&result.Products, "products")

	if result.Name != "" {
		return result, nil
	}
	return result, constants.CAT_ID_ERR
}

func DeleteCategoryByID(id uint) {
	cat := new(Category)
	cat.ID = id

	db := getConnectionDB()
	defer db.Close()
	db.Delete(&cat)
}
