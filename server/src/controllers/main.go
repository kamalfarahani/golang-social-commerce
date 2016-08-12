package controllers

import "github.com/kataras/iris"

func Register() {
	iris.Get("product/:id", getProduct)
	iris.Get("products/page/:pageNum", getProductsByPage)
	iris.Get("category", getAllCategories)
	iris.Get("category/:name", getCategory)
	iris.Get("category/:name/page/:pageNum",
		getCategoryProdcutsByPage)

	iris.Listen(":8585")
}

func addJsonHeader(context *iris.Context) {
	context.SetHeader("Content-Type", "application/json")
	context.SetHeader("Access-Control-Allow-Origin", "*")
	context.SetHeader("Access-Control-Allow-Headers",
		"Origin, X-Requested-With, Content-Type, Accept")
}
