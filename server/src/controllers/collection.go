package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/kataras/iris"

	"../converters"
	"../models"
)

func getCollection(context *iris.Context) {
	addAccessHeaders(context)

	name := context.Param("name")
	rawCol, err := models.GetCollectionByName(name)
	if err != nil {
		context.WriteString(err.Error())
		return
	}

	vmCol :=
		converters.ConvertCollectionToView(*rawCol)
	jsonCol, _ := json.Marshal(&vmCol)
	addJsonHeader(context)
	context.WriteString(string(jsonCol))
}

func getAllCollections(context *iris.Context) {
	addAccessHeaders(context)

	rawCols := models.GetAllCollections()
	vmCols :=
		converters.ConvertCollectionsToViews(rawCols)
	jsonCols, _ := json.Marshal(vmCols)
	addJsonHeader(context)
	context.WriteString(string(jsonCols))
}

func getCollectionsByPage(context *iris.Context) {
	addAccessHeaders(context)

	page, err := context.ParamInt("pageNum")
	if err != nil {
		context.NotFound()
		return
	}

	rawCols, err :=
		models.GetCollectionsByPage(uint(page))
	if err != nil {
		context.WriteString(err.Error())
		return
	}

	vmCols :=
		converters.ConvertCollectionsToViews(rawCols)
	jsonCols, _ := json.Marshal(vmCols)
	addJsonHeader(context)
	context.WriteString(string(jsonCols))
}

func getCollectionsPageCount(context *iris.Context) {
	addAccessHeaders(context)

	countStr :=
		strconv.Itoa(int(models.GetCollectionsPageCount()))
	context.WriteString(countStr)
}
