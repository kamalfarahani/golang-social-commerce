package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string `sql:"size:50"`
	Description string
	ImgUrl      string
	Price       int
}

func CreateProductTable() {
	db := getConnectionDB()
	defer db.Close()

	if !db.HasTable(&Product{}) {
		db.CreateTable(&Product{})
	}
}

func SaveProduct(name, description string) {
	db := getConnectionDB()
	defer db.Close()
	db.Save(
		&Product{Name: name,
			Description: description})
}

func GetProductsByPage(pageNum int) *[]Product {
	resultArr := new([]Product)
	offset := (pageNum - 1) * 10

	db := getConnectionDB()
	defer db.Close()
	db.Limit(10).Offset(offset).Find(resultArr)

	if len(*resultArr) > 0 {
		return resultArr
	}
	return nil
}

func GetProductByID(id int) *Product {
	result := new(Product)

	db := getConnectionDB()
	defer db.Close()
	db.First(result, id)

	if result.Name != "" || result.Description != "" {
		return result
	}
	return nil
}

func GetProductsByName(name string) *[]Product {
	resultArr := new([]Product)

	db := getConnectionDB()
	defer db.Close()
	db.Where("name = ?", name).Find(resultArr)

	if len(*resultArr) > 0 {
		return resultArr
	}
	return nil
}
