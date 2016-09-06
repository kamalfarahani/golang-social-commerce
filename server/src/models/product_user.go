package models

func GetUserLikedProductsByName(userID uint) []Product {
	user, _ := GetUserByID(userID)
	var products []Product

	db := getConnectionDB()
	defer db.Close()
	db.Model(&user).
		Association("LikedProducts").Find(&products)

	return products
}

func AddProductLikes(productID uint, userID uint) bool {
	product, productErr := GetProductByID(productID)
	user, userErr := GetUserByID(userID)

	if productErr != nil || userErr != nil {
		return false
	}

	db := getConnectionDB()
	defer db.Close()

	if IsProductInUserLikes(userID, productID) {
		product.Likes -= 1
		db.Model(user).
			Association("LikedProducts").Delete(product)
	} else {
		product.Likes += 1
		db.Model(user).
			Association("LikedProducts").Append(product)
	}
	db.Save(product)

	return true
}

func IsProductInUserLikes(userID uint, productID uint) bool {
	products :=
		GetUserLikedProductsByName(userID)

	for _, product := range products {
		if product.ID == productID {
			return true
		}
	}
	return false
}
