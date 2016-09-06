package controllers

import (
	"github.com/kataras/iris"

	"../models"
)

func registerUser(context *iris.Context) {
	nameStr :=
		string(context.FormValue("name"))
	passwordStr :=
		string(context.FormValue("password"))
	emailStr :=
		string(context.FormValue("email"))

	err := models.SaveUser(
		nameStr, emailStr, passwordStr)

	if err != nil {
		context.Write(err.Error())
		return
	}
	context.Write("Registered : %s", nameStr)
}

func loginUser(context *iris.Context) {
	nameStr :=
		string(context.FormValue("name"))
	passwordStr :=
		string(context.FormValue("password"))

	setUserSession(nameStr, passwordStr, context)
}

func setUserSession(name, password string, context *iris.Context) {
	user, err :=
		models.GetUserByNameAndPassword(name, password)
	if err == nil {
		context.Session().Set("user_id", user.ID)
		context.Write("welcome %s", name)
	} else {
		context.Write(err.Error())
	}
}

func isUserLogined(context *iris.Context) bool {
	if getSessionInt(context) != -1 {
		return true
	}
	return false
}

func getSessionInt(context *iris.Context) int {
	smap := context.Session().GetAll()

	if userID, ok := smap["user_id"].(uint); ok {
		return int(userID)
	}
	return -1
}
