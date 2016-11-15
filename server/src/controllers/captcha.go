package controllers

import (
	"image/png"

	"github.com/kataras/iris"

	"./captchaSystem"
)

func makeCaptchaImage(context *iris.Context) {
	context.SetHeader("Content-Type", "image/png")
	capImg := captchaSystem.MakeCaptcha(context)
	png.Encode(context.Response.BodyWriter(), capImg)
}
