package controllers

import (
	"github.com/kataras/iris"

	"../models"
)

func likeProduct(context *iris.Context) {
	addAccessHeaders(context)

	if !isUserLogined(context) {
		context.Write("Please login")
		return
	}
	userID := getSessionInt(context)

	productID, err := context.ParamInt("id")
	if err != nil {
		context.NotFound()
		return
	}

	if !models.AddProductLikes(uint(productID), uint(userID)) {
		context.Write("Data is wrong")
		return
	}

	context.Write("Done")
}
