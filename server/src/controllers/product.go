package controllers

import (
	"encoding/json"
	"github.com/kataras/iris"
	"kamal/server/social-commerce/server/src/converters"
	"kamal/server/social-commerce/server/src/models"
)

func getProduct(context *iris.Context) {
	id, err := context.ParamInt("id")

	if err == nil {
		rawProduct := models.GetProductByID(id)

		if rawProduct != nil {
			vmProduct :=
				converters.ConvertProductToView(*rawProduct)
			jsonProduct, _ := json.Marshal(&vmProduct)
			addJsonHeader(context)
			context.Write(string(jsonProduct))
		} else {
			context.NotFound()
		}
	} else {
		context.NotFound()
	}
}

func getProductsByPage(context *iris.Context) {
	page, err := context.ParamInt("pageNum")

	if err == nil {
		rawProducts := models.GetProductsByPage(page)

		if rawProducts != nil {
			vmProducts :=
				converters.ConvertProductsToViews(*rawProducts)
			jsonProduct, _ := json.Marshal(&vmProducts)
			addJsonHeader(context)
			context.Write(string(jsonProduct))
		} else {
			context.NotFound()
		}
	} else {
		context.NotFound()
	}
}
