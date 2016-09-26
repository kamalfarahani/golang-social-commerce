package models

import "../constants"

func SaveOrder(productID uint, userID uint, number uint) error {
	_, productErr := GetProductByID(productID)
	_, userErr := GetUserByID(userID)

	if productErr != nil || userErr != nil {
		return constants.WRONG_DATA_ERR
	}

	db := getConnectionDB()
	defer db.Close()
	db.Save(&Order{
		ProductID: productID,
		UserID:    userID,
		Number:    number,
	})

	return nil
}

func GetOrdersByUserName(userName string) ([]Order, error) {
	var orders []Order
	user, err := GetUserByName(userName)

	if err != nil {
		return orders, err
	}

	db := getConnectionDB()
	defer db.Close()
	db.Model(&user).
		Related(&orders, "Orders")

	return orders, nil
}
