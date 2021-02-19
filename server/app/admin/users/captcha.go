package users

import (
	"afkser/model/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type captcha struct {
	CaptchaID string `json:"captcha_id"`
	PicPath   string `json:"pic"`
}

// GenerateImgCaptcha 生成验证码
func GenerateImgCaptcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		response.Result(c, response.ServiceErr, err.Error())
	} else {
		response.ResultWithData(c, response.Success, "success", captcha{
			CaptchaID: id,
			PicPath:   b64s,
		})
	}
}
