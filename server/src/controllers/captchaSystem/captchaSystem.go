package captchaSystem

import (
	"time"

	"github.com/kamalmax/randString"
	"github.com/kataras/iris"
)

var cookieTime = make(map[string]time.Time)
var cookieCaptcha = make(map[string]string)

func init() {
	go captchaGarbageCollector()
}

func setCaptchaCookie(context *iris.Context, capStr string) {
	cookie := generateCaptchaCookieValue()
	context.SetCookieKV("captcha", cookie)

	cookieCaptcha[cookie] = capStr
	cookieTime[cookie] = time.Now()
}

func generateCaptchaCookieValue() string {
	exist := true
	var cookieValue string
	for exist {
		cookieValue = randString.RandomString(15)
		_, exist = cookieCaptcha[cookieValue]
	}

	return cookieValue
}

func captchaGarbageCollector() {
	for {
		time.Sleep(1 * time.Minute)
		for cookie, t := range cookieTime {
			if time.Since(t) >= 1*time.Minute {
				delete(cookieCaptcha, cookie)
				delete(cookieTime, cookie)
			}
		}
	}
}
