// middleware/captcha_middleware.go
package middleware

import (
	"Blog-Server/common/res"
	"Blog-Server/utils/email_store"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

type EmailVerifyMiddlewareRequest struct {
	EmailID   string `json:"emailID" binding:"required"`
	EmailCode string `json:"emailCode" binding:"required"`
}

func EmailVerifyMiddleware(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		res.FailWithMsg("获取请求体错误", c)
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	var cr EmailVerifyMiddlewareRequest
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		logrus.Errorf("邮箱验证失败 %s", err)
		res.FailWithMsg("邮箱验证失败", c)
		c.Abort()
		return
	}

	info, ok := email_store.Verify(cr.EmailID, cr.EmailCode)
	if !ok {
		res.FailWithMsg("邮箱验证失败", c)
		c.Abort()
		return
	}
	c.Set("email", info.Email)
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
}
