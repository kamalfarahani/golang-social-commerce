package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name     string `sql:"size:50;unique"`
	ImgUrl   string
	Products []Product `gorm:"many2many:category_product"`
}

func CreateCategoryTable() {
	db := getConnectionDB()
	defer db.Close()

	if !db.HasTable(&Category{}) {
		db.CreateTable(&Category{})
	}
}

func SaveCategory(name string, products *[]Product) {
	db := getConnectionDB()
	defer db.Close()
	db.Save(
		&Category{Name: name,
			Products: *products})
}

func GetAllCategories() *[]Category {
	resultArr := new([]Category)

	db := getConnectionDB()
	defer db.Close()
	db.Find(resultArr)

	return resultArr
}

func GetCategoryByName(name string) *Category {
	result := new(Category)

	db := getConnectionDB()
	defer db.Close()
	db.Where("name = ?", name).First(result)
	db.Model(result).
		Related(&result.Products, "products")

	if result.Name != "" {
		return result
	}
	return nil
}

func GetCategoryByID(id int) *Category {
	result := new(Category)

	db := getConnectionDB()
	defer db.Close()
	db.First(result, id)

	if result.Name != "" {
		return result
	}
	return nil
}
