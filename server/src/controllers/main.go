package controllers

import "github.com/kataras/iris"

func Register() {
	iris.Get("product/:id", getProduct)
	iris.Get("product/:id/like", likeProduct)
	iris.Get("products/pagecount", getProductsPageCount)
	iris.Get("products/page/:pageNum", getProductsByPage)

	iris.Get("category", getAllCategories)
	iris.Get("category/:name", getCategory)
	iris.Get("category/:name/page/:pageNum",
		getCategoryProdcutsByPage)

	iris.Get("collection", getAllCollections)
	iris.Get("collection/:name", getCollection)
	iris.Get("collections/pagecount", getCollectionsPageCount)
	iris.Get("collections/page/:pageNum", getCollectionsByPage)

	iris.Post("order/:productID", orderProduct)

	iris.Post("user/login", loginUser)
	iris.Post("user/register", registerUser)

	iris.Listen(":8585")
}
