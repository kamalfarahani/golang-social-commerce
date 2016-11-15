package captchaSystem

import (
	"github.com/afocus/captcha"
	"github.com/kataras/iris"

	"../../constants"
)

func MakeCaptcha(context *iris.Context) *captcha.Image {
	capt := captcha.New()
	capt.SetFont("../../assets/font/capfont.ttf")
	capImg, capStr := capt.Create(5, captcha.NUM)

	setCaptchaCookie(context, capStr)

	return capImg
}

func ValidateCaptcha(context *iris.Context) error {
	var err error
	cookie := context.GetCookie("captcha")
	capStr := context.FormValueString("captcha")
	if cookie == "" ||
		capStr == "" ||
		cookieCaptcha[cookie] != capStr {
		err = constants.WRONG_CAPTCHA_ERR
	}

	context.RemoveCookie("captcha")
	delete(cookieCaptcha, cookie)
	delete(cookieTime, cookie)

	return err
}
