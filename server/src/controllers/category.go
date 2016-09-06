package controllers

import (
	"encoding/json"

	"github.com/kataras/iris"

	"../converters"
	"../models"
)

func getCategory(context *iris.Context) {
	name := context.Param("name")

	rawCat, err := models.GetCategoryByName(name)
	if err != nil {
		context.Write(err.Error())
		return
	}

	vmCat :=
		converters.ConvertCategoryToView(*rawCat)
	jsonCat, _ := json.Marshal(&vmCat)
	addJsonHeader(context)
	context.Write(string(jsonCat))
}

func getAllCategories(context *iris.Context) {
	rawCats := models.GetAllCategories()
	vmCats :=
		converters.ConvertCategoriesToViews(rawCats)
	jsonCats, _ := json.Marshal(vmCats)
	addJsonHeader(context)
	context.Write(string(jsonCats))
}
