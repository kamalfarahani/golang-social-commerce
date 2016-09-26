package models

import "../constants"

func GetProductsByCategoryName(catName string) ([]Product, error) {
	var productsArr []Product
	cat, err := GetCategoryByName(catName)
	if err != nil {
		return productsArr, err
	}

	db := getConnectionDB()
	defer db.Close()
	db.Model(cat).
		Related(&productsArr, "products")

	if len(productsArr) > 0 {
		return productsArr, nil
	}
	return productsArr, constants.CAT_NO_PRODUCT_ERR
}

func GetProductsByCategoryID(catID uint) ([]Product, error) {
	var productsArr []Product
	cat, err := GetCategoryByID(catID)
	if err != nil {
		return productsArr, err
	}

	db := getConnectionDB()
	defer db.Close()
	db.Model(cat).
		Related(&productsArr, "products")

	if len(productsArr) > 0 {
		return productsArr, nil
	}
	return productsArr, constants.CAT_NO_PRODUCT_ERR
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
	return productsArr, constants.CAT_END_PRODUCT_ERR
}

func GetCategoryByProductID(productID uint) (*Category, error) {
	product, err := GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	return GetCategoryByID(product.CatID)
}
