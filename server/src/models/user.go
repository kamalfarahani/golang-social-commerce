package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Password      string
	Name          string    `sql:"unique;unique_index"`
	Email         string    `sql:"unique;unique_index"`
	Orders        []Order   `gorm:"ForeignKey:UserID"`
	LikedProducts []Product `gorm:"many2many:user_product_likes;"`
}

func CreateUserTable() {
	db := getConnectionDB()
	defer db.Close()

	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
}

func SaveUser(name, email, password string) error {
	if err := userValidation(name, email, password); err != nil {
		return err
	}

	hashedPass := hasheString(password)
	user := User{
		Name:     name,
		Email:    email,
		Password: hashedPass,
	}

	db := getConnectionDB()
	defer db.Close()
	err := db.Save(&user).Error

	return err
}

func GetUserByID(id uint) (*User, error) {
	result := new(User)

	db := getConnectionDB()
	defer db.Close()
	db.First(result, id)

	if result.Name != "" {
		return result, nil
	}
	return result, errors.New("id is wrong")
}

func GetUserByNameAndPassword(name, password string) (*User, error) {
	result := new(User)
	hashedPass := hasheString(password)

	db := getConnectionDB()
	defer db.Close()
	db.Where("name = ? AND password = ?", name, hashedPass).
		First(result)

	if result.Name != "" {
		return result, nil
	}
	return result, errors.New("name or password is wrong")
}

func GetUserByName(name string) (*User, error) {
	result := new(User)

	db := getConnectionDB()
	defer db.Close()
	db.Where("name = ?", name).
		First(result)

	if result.Name != "" {
		return result, nil
	}
	return result, errors.New("name is wrong")
}

func hasheString(str string) string {
	hashByte := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hashByte[:])
}

func userValidation(name, email, password string) error {
	if len(name) < 3 {
		return errors.New("Name is too short")
	}

	if len(password) < 6 {
		return errors.New("Password is too short")
	}

	if !govalidator.IsEmail(email) {
		return errors.New("The email is invalid")
	}

	return nil
}
