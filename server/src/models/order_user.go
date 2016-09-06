package models

import "errors"

func SaveOrder(productID uint, userID uint, number uint) error {
	_, productErr := GetProductByID(productID)
	_, userErr := GetUserByID(userID)

	if productErr != nil || userErr != nil {
		return errors.New("wrong data")
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
