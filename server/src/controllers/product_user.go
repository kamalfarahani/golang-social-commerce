package controllers

import (
	"github.com/kataras/iris"

	"../models"
)

func likeProduct(context *iris.Context) {
	addAccessHeaders(context)

	userID, err := getSessionInt(context)
	if err != nil {
		context.WriteString(err.Error())
		return
	}

	productID, err := context.ParamInt("id")
	if err != nil {
		context.NotFound()
		return
	}

	if !models.AddProductLikes(uint(productID), uint(userID)) {
		context.WriteString("Data is wrong")
		return
	}

	context.WriteString("Done")
}
