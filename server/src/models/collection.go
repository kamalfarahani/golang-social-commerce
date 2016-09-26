package models

import (
	"github.com/jinzhu/gorm"

	"../constants"
)

type Collection struct {
	gorm.Model
	ImgUrl   string
	Name     string    `sql:"size:50;unique;unique_index"`
	Products []Product `gorm:"many2many:collection_product;"`
}

func CreateCollectionTable() {
	db := getConnectionDB()
	defer db.Close()

	if !db.HasTable(&Collection{}) {
		db.CreateTable(&Collection{})
	}
}

func SaveCollection(name string, products []Product) {
	db := getConnectionDB()
	defer db.Close()
	db.Save(
		&Collection{
			Name:     name,
			Products: products,
		})
}

func GetAllCollections() []Collection {
	var colArr []Collection

	db := getConnectionDB()
	defer db.Close()
	db.Find(&colArr)

	return colArr
}

func GetCollectionsByPage(pageNum uint) ([]Collection, error) {
	if pageNum > GetCollectionsPageCount() {
		return nil, constants.END_PAGE_ERR
	}

	var collectionsArr []Collection
	offset := (pageNum - 1) * 10

	db := getConnectionDB()
	defer db.Close()
	db.Limit(10).
		Offset(offset).Find(&collectionsArr)

	return collectionsArr, nil
}

func GetCollectionByName(name string) (*Collection, error) {
	result := new(Collection)

	db := getConnectionDB()
	defer db.Close()
	db.Where("name = ?", name).First(result)
	db.Model(result).
		Related(&result.Products, "products")

	if result.Name != "" {
		return result, nil
	}
	return result, constants.COL_NAME_ERR
}

func GetCollectionByID(id uint) (*Collection, error) {
	result := new(Collection)

	db := getConnectionDB()
	defer db.Close()
	db.First(result, id)
	db.Model(result).
		Related(&result.Products, "products")

	if result.Name != "" {
		return result, nil
	}
	return result, constants.COL_ID_ERR
}

func GetCollectionsPageCount() uint {
	var collectionsCount uint

	db := getConnectionDB()
	defer db.Close()
	db.Model(&Collection{}).Count(&collectionsCount)

	pageCountInt :=
		(collectionsCount / 10)
	pageCountFloat :=
		(float32(collectionsCount) / 10.0)

	if float32(pageCountInt) < pageCountFloat {
		return (pageCountInt + 1)
	}
	return pageCountInt
}
