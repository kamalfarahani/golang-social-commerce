package controllers

import (
	"github.com/kataras/iris"
	"strconv"

	"../models"
)

func orderProduct(context *iris.Context) {
	addAccessHeaders(context)

	if !isUserLogined(context) {
		context.Write("Please login!")
		return
	}
	userID := getSessionInt(context)

	productID, err := context.ParamInt("productID")
	if err != nil {
		context.NotFound()
		return
	}

	orderNum, err := strconv.Atoi(
		string(context.FormValue("order_number")))
	if err != nil {
		context.Write(err.Error())
		return
	}

	err = models.SaveOrder(
		uint(productID), uint(userID), uint(orderNum))

	if err != nil {
		context.Write(err.Error())
		return
	}
	context.Write("Submited")
}
