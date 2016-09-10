package controllers

import "github.com/kataras/iris"

func addAccessHeaders(context *iris.Context) {
	context.SetHeader("Access-Control-Allow-Origin", "*")
	context.SetHeader("Access-Control-Allow-Headers",
		"Origin, X-Requested-With, Content-Type, Accept")
}

func addJsonHeader(context *iris.Context) {
	context.SetHeader("Content-Type", "application/json")
}
