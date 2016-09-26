package models

import (
	"github.com/jinzhu/gorm"

	"../constants"
)

type Product struct {
	gorm.Model
	Name        string `sql:"size:50"`
	Description string
	ImgUrl      string
	Price       int
	CatID       uint
	Likes       uint
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
		&Product{
			Name:        name,
			Description: description,
		})
}

func GetProductByID(id uint) (*Product, error) {
	result := new(Product)

	db := getConnectionDB()
	defer db.Close()
	db.First(result, id)

	if result.Name != "" {
		return result, nil
	}
	return result, constants.PRODUCT_ID_ERR
}

func GetProductsByPage(pageNum uint) ([]Product, error) {
	if pageNum > GetProductsPageCount() {
		return nil, constants.END_PAGE_ERR
	}

	var productsArr []Product
	offset := (pageNum - 1) * 10

	db := getConnectionDB()
	defer db.Close()
	db.Limit(10).
		Offset(offset).Find(&productsArr)

	return productsArr, nil
}

func GetProductsByName(name string) ([]Product, error) {
	var productsArr []Product

	db := getConnectionDB()
	defer db.Close()
	db.Where("name = ?", name).Find(&productsArr)

	if len(productsArr) > 0 {
		return productsArr, nil
	} else {
		return productsArr, constants.PRODUCT_NAME_ERR
	}
}

func GetProductsPageCount() uint {
	var productsCount uint

	db := getConnectionDB()
	defer db.Close()
	db.Model(&Product{}).Count(&productsCount)

	pageCountInt :=
		(productsCount / 10)
	pageCountFloat :=
		(float32(productsCount) / 10.0)

	if float32(pageCountInt) < pageCountFloat {
		return (pageCountInt + 1)
	}
	return pageCountInt
}
