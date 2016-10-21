package controllers

import (
	"image/png"
	"net/http"

	"./captchaSystem"
)

func makeCaptchaImage(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "image/png")
	capImg := captchaSystem.MakeCaptcha(req)
	png.Encode(rw, capImg)
}
