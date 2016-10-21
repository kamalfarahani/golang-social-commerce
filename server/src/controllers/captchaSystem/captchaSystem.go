package captchaSystem

import (
	"net/http"
	"strings"
	"time"

	"github.com/afocus/captcha"
	"github.com/kataras/iris"

	"../../constants"
)

var ipTime = make(map[string]time.Time)
var ipCaptcha = make(map[string]string)

func init() {
	go deleteOldCaptchas()
}

func MakeCaptcha(req *http.Request) *captcha.Image {
	capt := captcha.New()
	capt.SetFont("../../assets/font/capfont.ttf")
	capImg, capStr := capt.Create(5, captcha.NUM)

	remoteAddr := req.RemoteAddr
	ip := remoteAddr[:strings.Index(remoteAddr, ":")]
	ipCaptcha[ip] = capStr
	ipTime[ip] = time.Now()

	return capImg
}

func ValidateCaptchaByIP(context *iris.Context) error {
	ip := context.RequestIP()
	capStr := context.FormValueString("captcha")
	if capStr != "" && ipCaptcha[ip] == capStr {
		delete(ipCaptcha, ip)
		delete(ipTime, ip)

		return nil
	}

	delete(ipCaptcha, ip)
	delete(ipTime, ip)

	return constants.WRONG_CAPTCHA_ERR
}

func deleteOldCaptchas() {
	for {
		time.Sleep(1 * time.Minute)
		for ip, t := range ipTime {
			if time.Since(t) >= 1*time.Minute {
				delete(ipCaptcha, ip)
				delete(ipTime, ip)
			}
		}
	}
}
