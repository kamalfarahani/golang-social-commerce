package controllers

import (
	"encoding/json"
	"github.com/kataras/iris"
	"kamal/server/social-commerce/server/src/converters"
	"kamal/server/social-commerce/server/src/models"
)

func getCategory(context *iris.Context) {
	name := context.Param("name")
	rawCat := models.GetCategoryByName(name)

	if rawCat != nil {
		vmCat :=
			converters.ConvertCategoryToView(*rawCat)
		jsonCat, _ := json.Marshal(&vmCat)
		addJsonHeader(context)
		context.Write(string(jsonCat))
	} else {
		context.NotFound()
	}
}

func getAllCategories(context *iris.Context) {
	rawCats := models.GetAllCategories()
	vmCats :=
		converters.ConvertCategoriesToViews(*rawCats)
	jsonCats, _ := json.Marshal(&vmCats)
	addJsonHeader(context)
	context.Write(string(jsonCats))
}
