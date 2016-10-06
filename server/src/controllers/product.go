package controllers

import (
	"strconv"

	"github.com/kataras/iris"

	"../converters"
	"../models"
)

func getProduct(context *iris.Context) {
	addAccessHeaders(context)

	id, err := context.ParamInt("id")
	if err != nil {
		context.NotFound()
		return
	}

	rawProduct, err :=
		models.GetProductByID(uint(id))
	if err != nil {
		context.WriteString(err.Error())
		return
	}

	vmProduct :=
		converters.ConvertProductToView(*rawProduct)
	jsonProduct := jsonStr(&vmProduct)
	addJsonHeader(context)
	context.WriteString(jsonProduct)
}

func getProductsByPage(context *iris.Context) {
	addAccessHeaders(context)

	page, err := context.ParamInt("pageNum")
	if err != nil {
		context.NotFound()
		return
	}

	rawProducts, err :=
		models.GetProductsByPage(uint(page))
	if err != nil {
		context.WriteString(err.Error())
		return
	}

	vmProducts :=
		converters.ConvertProductsToViews(rawProducts)
	jsonProduct := jsonStr(vmProducts)
	addJsonHeader(context)
	context.WriteString(jsonProduct)
}

func getProductsPageCount(context *iris.Context) {
	addAccessHeaders(context)

	countStr :=
		strconv.Itoa(int(models.GetProductsPageCount()))
	context.WriteString(countStr)
}
