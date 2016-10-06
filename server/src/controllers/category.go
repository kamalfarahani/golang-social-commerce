package controllers

import (
	"github.com/kataras/iris"

	"../converters"
	"../models"
)

func getCategory(context *iris.Context) {
	addAccessHeaders(context)

	name := context.Param("name")
	rawCat, err := models.GetCategoryByName(name)
	if err != nil {
		context.WriteString(err.Error())
		return
	}

	vmCat :=
		converters.ConvertCategoryToView(*rawCat)
	jsonCat := jsonStr(&vmCat)
	addJsonHeader(context)
	context.WriteString(jsonCat)
}

func getAllCategories(context *iris.Context) {
	addAccessHeaders(context)

	rawCats := models.GetAllCategories()
	vmCats :=
		converters.ConvertCategoriesToViews(rawCats)
	jsonCats := jsonStr(vmCats)
	addJsonHeader(context)
	context.WriteString(jsonCats)
}
