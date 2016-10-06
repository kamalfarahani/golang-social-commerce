package controllers

import (
	"github.com/kataras/iris"

	"../converters"
	"../models"
)

func getCategoryProdcutsByPage(context *iris.Context) {
	addAccessHeaders(context)

	name := context.Param("name")
	page, err := context.ParamInt("pageNum")
	if err != nil {
		context.NotFound()
		return
	}

	rawProducts, err :=
		models.GetCategoryProductsByPage(name, uint(page))
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
