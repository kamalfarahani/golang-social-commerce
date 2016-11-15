package controllers

import (
	"github.com/kataras/iris"

	"../constants"
	"../models"
	"./captchaSystem"
)

func registerUser(context *iris.Context) {
	addAccessHeaders(context)

	err :=
		captchaSystem.ValidateCaptcha(context)
	if err != nil {
		context.WriteString(err.Error())
		return
	}

	nameStr :=
		context.FormValueString("name")
	emailStr :=
		context.FormValueString("email")
	passwordStr :=
		context.FormValueString("password")

	err =
		models.SaveUser(nameStr, emailStr, passwordStr)
	if err != nil {
		context.WriteString(err.Error())
		return
	}

	context.Write("Registered : %s", nameStr)
}

func loginUser(context *iris.Context) {
	addAccessHeaders(context)

	err :=
		captchaSystem.ValidateCaptcha(context)
	if err != nil {
		context.WriteString(err.Error())
		return
	}

	nameStr :=
		context.FormValueString("name")
	passwordStr :=
		context.FormValueString("password")

	setUserSession(nameStr, passwordStr, context)
}

func setUserSession(name, password string, context *iris.Context) {
	user, err :=
		models.GetUserByNameAndPassword(name, password)
	if err != nil {
		context.WriteString(err.Error())
		return
	}

	context.Session().Set("user_id", user.ID)
	context.Write("welcome %s", name)
}

func getSessionInt(context *iris.Context) (int, error) {
	smap := context.Session().GetAll()

	if userID, ok := smap["user_id"].(uint); ok {
		return int(userID), nil
	}
	return -1, constants.USER_NOT_LOGIN
}
