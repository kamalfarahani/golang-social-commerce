package models

import "../constants"

func GetProductsByCollectionName(colName string) ([]Product, error) {
	var productsArr []Product
	col, err := GetCollectionByName(colName)
	if err != nil {
		return productsArr, err
	}

	db := getConnectionDB()
	defer db.Close()
	db.Model(col).
		Related(&productsArr, "products")

	if len(productsArr) > 0 {
		return productsArr, nil
	}
	return productsArr, constants.COL_NO_PRODUCT_ERR
}

func GetProductsByCollectionID(colID uint) ([]Product, error) {
	var productsArr []Product
	col, err := GetCollectionByID(colID)
	if err != nil {
		return productsArr, err
	}

	db := getConnectionDB()
	defer db.Close()
	db.Model(col).
		Related(&productsArr, "products")

	if len(productsArr) > 0 {
		return productsArr, nil
	}
	return productsArr, constants.COL_NO_PRODUCT_ERR
}

// this function has wrong result
func GetCollectionsByProductID(productID uint) (*Collection, error) {
	result := new(Collection)

	product, err := GetProductByID(productID)
	if err != nil {
		return result, err
	}

	db := getConnectionDB()
	defer db.Close()
	db.Model(product).
		Related(result)

	return result, nil
}
