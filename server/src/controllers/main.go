package controllers

import "github.com/kataras/iris"

func Register() {
	iris.Get("product/:id", getProduct)
	iris.Get("product/:id/like", likeProduct)
	iris.Get("products/page/:pageNum", getProductsByPage)
	iris.Get("products/pagecount", getProductsPageCount)

	iris.Get("category", getAllCategories)
	iris.Get("category/:name", getCategory)
	iris.Get("category/:name/page/:pageNum",
		getCategoryProdcutsByPage)

	iris.Post("order/:productID", orderProduct)

	iris.Post("user/register", registerUser)
	iris.Post("user/login", loginUser)

	iris.Listen(":8585")
}
