package models

func GetProductsByCategoryName(catName string) *[]Product {
	resultArr := new([]Product)
	cat := new(Category)

	db := getConnectionDB()
	defer db.Close()
	db.Where("name = ?", catName).First(cat)
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

func GetCategoryProductsByPage(catName string, pageNum int) *[]Product {
	resultArr := new([]Product)
	cat := new(Category)
	offset := (pageNum - 1) * 10

	db := getConnectionDB()
	defer db.Close()
	db.Limit(10).Offset(offset).
		Where("name = ?", catName).First(cat)
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
