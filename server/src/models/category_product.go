package models

func GetProductsByCategoryName(name string) *[]Product {
	resultArr := new([]Product)
	cat := new(Category)

	db := getConnectionDB()
	defer db.Close()
	db.Where("name = ?", name).First(cat)
	db.Model(cat).
		Related(resultArr, "products")

	if len(*resultArr) > 0 {
		return resultArr
	} else {
		return nil
	}
}

func GetProductsByCategoryID(id int) *[]Product {
	resultArr := new([]Product)
	cat := new(Category)

	db := getConnectionDB()
	defer db.Close()
	db.First(cat, id)
	db.Model(cat).
		Related(resultArr, "products")

	if len(*resultArr) > 0 {
		return resultArr
	} else {
		return nil
	}
}

func GetCategoriesByProductID(id int) *[]Category {
	resultArr := new([]Category)
	product := new(Product)

	db := getConnectionDB()
	defer db.Close()
	db.First(product, id)
	db.Model(product).
		Related(resultArr)

	if len(*resultArr) > 0 {
		return resultArr
	} else {
		return nil
	}
}
