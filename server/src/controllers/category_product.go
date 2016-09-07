package controllers

import (
	"encoding/json"

	"github.com/kataras/iris"

	"../converters"
	"../models"
)

func getCategoryProdcutsByPage(context *iris.Context) {
	name := context.Param("name")
	page, err := context.ParamInt("pageNum")
	if err != nil {
		context.NotFound()
		return
	}

	rawProducts, err :=
		models.GetCategoryProductsByPage(name, uint(page))
	if err != nil {
		context.Write(err.Error())
		return
	}

	vmProducts :=
		converters.ConvertProductsToViews(rawProducts)
	jsonProduct, _ := json.Marshal(vmProducts)
	addJsonHeader(context)
	context.Write(string(jsonProduct))
}
