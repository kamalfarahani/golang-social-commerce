package models

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

func CreateUserTable() {
	db := getConnectionDB()
	defer db.Close()

	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
}

func SaveUser(name, password string) {
	hashedPass := haseString(password)

	user := User{
		Name:     name,
		Password: hashedPass,
	}

	db := getConnectionDB()
	defer db.Close()
	db.Save(&user)
}

func GetUser(name, password string) *User {
	result := new(User)
	hasedPass := haseString(password)

	db := getConnectionDB()
	defer db.Close()
	db.Where("name = ? AND password = ?", name, hasedPass).
		First(result)

	if result.Name != "" {
		return result
	}
	return nil
}

func haseString(str string) string {
	hashByte := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hashByte[:])
}
