package models

import "errors"

func GetProductsByCategoryName(catName string) ([]Product, error) {
	var productsArr []Product
	cat, err := GetCategoryByName(catName)
	if err != nil {
		return productsArr,
			errors.New("invalid category name")
	}

	db := getConnectionDB()
	defer db.Close()
	db.Model(cat).
		Related(&productsArr, "products")

	if len(productsArr) > 0 {
		return productsArr, nil
	}
	return productsArr, errors.New("Category has no product")
}

func GetProductsByCategoryID(id uint) ([]Product, error) {
	var productsArr []Product
	cat, err := GetCategoryByID(id)
	if err != nil {
		return productsArr,
			errors.New("invalid category id")
	}

	db := getConnectionDB()
	defer db.Close()
	db.Model(cat).
		Related(&productsArr, "products")

	if len(productsArr) > 0 {
		return productsArr, nil
	}
	return productsArr, errors.New("Category has no product")
}

func GetCategoryProductsByPage(catName string, pageNum uint) ([]Product, error) {
	var productsArr []Product
	cat, _ := GetCategoryByName(catName)
	offset := (pageNum - 1) * 10

	db := getConnectionDB()
	defer db.Close()
	db.Model(cat).Limit(10).Offset(offset).
		Related(&productsArr, "products")

	if len(productsArr) > 0 {
		return productsArr, nil
	}
	return productsArr, errors.New("No more products")
}

func GetCategoryByProductID(id uint) (*Category, error) {
	product, err := GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return GetCategoryByID(product.CatID)
}
