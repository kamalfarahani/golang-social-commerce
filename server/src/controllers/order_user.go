package controllers

import (
	"github.com/kataras/iris"
	"strconv"

	"../models"
)

func orderProduct(context *iris.Context) {
	addAccessHeaders(context)

	userID, err := getSessionInt(context)
	if err != nil {
		context.WriteString(err.Error())
		return
	}

	productID, err := context.ParamInt("productID")
	if err != nil {
		context.NotFound()
		return
	}

	orderNum, err := strconv.Atoi(
		string(context.FormValue("order_number")))
	if err != nil {
		context.WriteString(err.Error())
		return
	}

	err = models.SaveOrder(
		uint(productID), uint(userID), uint(orderNum))

	if err != nil {
		context.WriteString(err.Error())
		return
	}
	context.WriteString("Submited")
}
