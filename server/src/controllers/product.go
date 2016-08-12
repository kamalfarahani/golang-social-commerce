package controllers

import (
	"encoding/json"
	"github.com/kataras/iris"
	"kamal/server/social-commerce/server/src/converters"
	"kamal/server/social-commerce/server/src/models"
)

func getProduct(context *iris.Context) {
	id, err := context.ParamInt("id")
	if err != nil {
		context.NotFound()
		return
	}

	rawProduct := models.GetProductByID(id)
	if rawProduct == nil {
		context.NotFound()
		return
	}

	vmProduct :=
		converters.ConvertProductToView(*rawProduct)
	jsonProduct, _ := json.Marshal(&vmProduct)
	addJsonHeader(context)
	context.Write(string(jsonProduct))
}

func getProductsByPage(context *iris.Context) {
	page, err := context.ParamInt("pageNum")
	if err != nil {
		context.NotFound()
		return
	}

	rawProducts := models.GetProductsByPage(page)
	if rawProducts == nil {
		context.NotFound()
		return
	}

	vmProducts :=
		converters.ConvertProductsToViews(*rawProducts)
	jsonProduct, _ := json.Marshal(&vmProducts)
	addJsonHeader(context)
	context.Write(string(jsonProduct))
}
